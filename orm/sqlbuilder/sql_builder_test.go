package sqlbuilder

import (
	"testing"
	"time"

	"github.com/ezbuy/redis-orm/example/model"
)

func TestSQLBuilder(t *testing.T) {

	testDate := time.Date(2017, 5, 27, 11, 20, 33, 0, time.Local)

	cases := []struct {
		b                  Builder
		mysqlOut, mssqlOut string
	}{
		{
			And(
				Eq("OrderId", []int64{
					11864555,
					11864554,
					11864553,
					11864552,
					11864551,
					11864550,
					11864549,
					11864548,
				}),
				Eq("PurchaseType", "Ezbuy"),
				Neq("PoPlaceDate", nil),
				Lt("OrderDate", testDate),
			),
			"(`OrderId` IN (11864555,11864554,11864553,11864552,11864551,11864550,11864549,11864548)) AND (`PurchaseType` = 'Ezbuy') AND (`PoPlaceDate` IS NOT NULL) AND (`OrderDate` < '2017-05-27 11:20:33.000000')",
			// select top 5 * from [order] where ([OrderId] IN (11864555,11864554,11864553,11864552,11864551,11864550,11864549,11864548)) AND ([PurchaseType] = N'Ezbuy') AND ([PoPlaceDate] IS NOT NULL) AND ([OrderDate] < N'2017-05-27 11:15:49.723')
			"([OrderId] IN (11864555,11864554,11864553,11864552,11864551,11864550,11864549,11864548)) AND ([PurchaseType] = N'Ezbuy') AND ([PoPlaceDate] IS NOT NULL) AND ([OrderDate] < N'2017-05-27 11:20:33.000')",
		},
		{
			Set().Add("ShipperName", "顺丰快递").Add("TrackingNo", "123223323423").Add("SyncDate", testDate),
			"`ShipperName` = '顺丰快递', `TrackingNo` = '123223323423', `SyncDate` = '2017-05-27 11:20:33.000000'",
			// update OrderTracking set [ShipperName] = N'顺丰快递', [TrackingNo] = N'123223323423', [SyncDate] = N'2017-05-27 11:20:33.000' where OrderTrackingId = 7739010;
			"[ShipperName] = N'顺丰快递', [TrackingNo] = N'123223323423', [SyncDate] = N'2017-05-27 11:20:33.000'",
		},
		{
			And(
				Eq(model.BlogColumns.Id, 1),
			),
			"`id` = 1",
			"[id] = 1",
		},
	}

	for i, c := range cases {
		mysqlOut := MySQL.MustBuild(c.b)

		mssqlOut := MSSQL.MustBuild(c.b)

		if c.mysqlOut != mysqlOut {
			t.Errorf("#%d [mysql] expected %q, got %q", i+1, c.mysqlOut, mysqlOut)
		}

		if c.mssqlOut != mssqlOut {
			t.Errorf("#%d [mssql] expected %q, got %q", i+1, c.mssqlOut, mssqlOut)
		}
	}
}
