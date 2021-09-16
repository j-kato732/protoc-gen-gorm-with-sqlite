package action

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	pb "grpc_gateway_sample/proto"
)

const (
	db_path = "./db/test.db"
)

var periods_orm []pb.PeriodORM

func GetPeriod(ctx context.Context) ([]*pb.Period, error) {
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	con, err := db.DB()
	defer con.Close()

	// SELECT * FROM Period;
	if err := db.Find(&periods_orm).Error; err != nil {
		return nil, err
	}

	var periods_pb []*pb.Period
	for _, period := range periods_orm {
		result, _ := period.ToPB(ctx)
		periods_pb = append(periods_pb, &result)
	}

	return periods_pb, nil
}
