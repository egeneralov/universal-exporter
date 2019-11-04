# universal-exporter

Prometheus exporter. May be for all.

I could not find an exporter for yota, I had to write my own. In the process made it universal. The default url contains the path to yota metrics.

To start the exporter, do not forget to specify `-mode exporter`. You can view the prometheus metrics using `-mode prometheus`. The default behavior is `-mode json`.

If you have something that gives metrics with a well-known separator, you can try feeding them a promethe through this utility.

#### Initial data:

    3GPP.SINR=2
    3GPP.RSSI=-66
    3GPP.RSRP=-97
    3GPP.RSRQ=-9
    3GPP.MCC=250

#### Will be generated for prometheus:

    # HELP s3gpp_sinr s3gpp_sinr 
    # TYPE s3gpp_sinr gauge
    s3gpp_sinr 2
    # HELP s3gpp_rssi s3gpp_rssi 
    # TYPE s3gpp_rssi gauge
    s3gpp_rssi -66
    # HELP s3gpp_rsrp s3gpp_rsrp 
    # TYPE s3gpp_rsrp gauge
    s3gpp_rsrp -97
    # HELP s3gpp_rsrq s3gpp_rsrq 
    # TYPE s3gpp_rsrq gauge
    s3gpp_rsrq -9
    # HELP s3gpp_mcc s3gpp_mcc 
    # TYPE s3gpp_mcc gauge
    s3gpp_mcc 250

#### Example json:

    {
      "3gpp-sinr": 1,
      "3gpp-rssi": -66,
      "3gpp-rsrp": -97,
      "3gpp-rsrq": -8,
      "3gpp-mcc": 250,
    }
