package v2

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_OrderRegisterStatus(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.OrderRegister(timedCtx, nil)
	require.Error(t, err)
	require.Nil(t, resp)

	resp, err = c.OrderRegister(timedCtx, &OrderRegisterRequest{
		Type:         0,
		Number:       "123",
		Comment:      "test",
		TariffCode:   62,
		FromLocation: OrderLocation{Code: 44, Address: "qwe"},
		ToLocation:   OrderLocation{Code: 287, Address: "qwe"},
		Sender: OrderSenderRecipient{
			Name:    "test",
			Company: "test",
			Email:   "test@test.com",
		},
		Recipient: OrderSenderRecipient{
			Name: "test",
			Phones: []OrderPhone{
				{Number: "123"},
			},
		},
		Packages: []OrderPackage{
			{
				Number:  "test",
				Weight:  1,
				Comment: "test",
				Items: []OrderPackageItem{
					{
						Name:    "test",
						WareKey: "test",
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Greater(t, len(resp.Requests), 0)

	// @todo ticket SD-735298
	fmt.Printf("\n\n!!!@@@ %+v\n\n", resp)

	statusResp, err := c.OrderStatus(ctx, resp.Entity.Uuid)
	require.NoError(t, err)
	require.Equal(t, statusResp.Entity.Comment, "test")
}
