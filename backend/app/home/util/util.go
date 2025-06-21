// Date:   Sat Jun 21 23:18:15 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func WithRPCClient[T any](
	addr string, 
	newClient func(conn grpc.ClientConnInterface) T, 
	callback func(ctx context.Context, client T),
) error {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	callback(ctx, newClient(conn))
	return nil
}
