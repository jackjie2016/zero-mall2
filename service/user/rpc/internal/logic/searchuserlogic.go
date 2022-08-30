package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/global"
	Grommodel "zero-mal/service/user/model/gorm"

	"github.com/zeromicro/go-zero/core/logx"

	"zero-mal/common/tool"
	"zero-mal/service/user/rpc/internal/svc"
	pb "zero-mal/service/user/rpc/user_pb"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *pb.SearchUserReq) (*pb.UserListResponse, error) {
	// todo: add your logic here and delete this line
	//关键词搜索、查询新品、查询热门商品、通过价格区间筛选， 通过商品分类筛选
	var users []Grommodel.User
	var userModel Grommodel.User

	_ = copier.Copy(&userModel, in)
	fmt.Printf("%+v", userModel)
	fmt.Printf("%#v", userModel)
	localDB := global.DB.Model(&Grommodel.User{})

	if in.Id > 0 {
		localDB = localDB.Where("id", in.Id)
	}
	if in.Mobile != "" {
		localDB = localDB.Where("mobile", in.Mobile)
	}

	if in.Gender != "" {
		localDB = localDB.Where(Grommodel.User{Gender: in.Gender})
	}
	if in.NickName != "" {
		localDB = localDB.Where("nick_name Like ?", "%"+in.NickName+"%")
		//多字段查询
		//q = q.Must(elastic.NewMultiMatchQuery(req.KeyWords, "name", "goods_brief"))
	}

	result := localDB.Scopes(tool.Paginate(int(in.Page), int(in.PageSize))).Find(&users)
	//查询没有错误
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "异常")
	}
	if result.RowsAffected == 0 {
		return &pb.UserListResponse{}, nil
	}

	rsp := &pb.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	userInfoRsp := pb.UserInfoResponse{}
	for _, user := range users {
		_ = copier.Copy(&userInfoRsp, user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	fmt.Println(rsp)
	return rsp, nil
}
