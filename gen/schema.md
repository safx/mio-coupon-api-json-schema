
## Coupon
クーポン状態の照会と切り替え

### Attributes
<table>
  <tr>
    <th>Name</th>
    <th>Type</th>
    <th>Description</th>
    <th>Example</th>
  </tr>
  <tr>
    <td><strong>couponInfo/coupon</strong></td>
    <td><em>array</em></td>
    <td></td>
    <td><code>[{"volume"=>42, "expire"=>"201403", "type"=>"bundle"}]</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hddServiceCode</strong></td>
    <td><em>string</em></td>
    <td>hddサービスコード</td>
    <td><code>"hdd01234567"</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/coupon</strong></td>
    <td><em>array</em></td>
    <td></td>
    <td><code>[{"volume"=>42, "expire"=>"201403", "type"=>"bundle"}]</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/couponUse</strong></td>
    <td><em>boolean</em></td>
    <td>クーポン使用中（クーポンON）の場合に true, それ以外は false</td>
    <td><code>false</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/hdoServiceCode</strong></td>
    <td><em>string</em></td>
    <td>hdoサービスコード</td>
    <td><code>"hdo01234567"</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/iccid</strong></td>
    <td><em>string</em></td>
    <td>回線のICCID</td>
    <td><code>"DN000011112222"</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/number</strong></td>
    <td><em>string</em></td>
    <td>回線の電話番号</td>
    <td><code>"08011112222"</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/regulation</strong></td>
    <td><em>boolean</em></td>
    <td>通信規制中の場合に true, それ以外は false</td>
    <td><code>false</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/sms</strong></td>
    <td><em>boolean</em></td>
    <td>SMS機能付きの場合に true, それ以外は false（音声通話機能付きSIMカードの場合も true が返ります）</td>
    <td><code>false</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/voice</strong></td>
    <td><em>boolean</em></td>
    <td>音声通話機能付きの場合に true, それ以外は false</td>
    <td><code>false</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/plan</strong></td>
    <td><em>string</em></td>
    <td>プラン名<br/><b>one of:</b><code>"Family Share"</code> or <code>"Minimum Start"</code> or <code>"Light Start"</code></td>
    <td><code>"Family Share"</code></td>
  </tr>
  <tr>
    <td><strong>returnCode</strong></td>
    <td><em>string</em></td>
    <td>リターンコード</td>
    <td><code>"OK"</code></td>
  </tr>
</table>

### Coupon List
クーポン残量とクーポンのON/OFF状態を照会します

```
GET /coupon/
```


#### Curl Example
```bash
$ curl -n -X GET https://api.iijmio.jp/mobile/d/v1/coupon/
```


#### Response Example
```
HTTP/1.1 200 OK
```
```json
[
  {
    "returnCode": "OK",
    "couponInfo": [
      {
        "hddServiceCode": "hdd01234567",
        "plan": "Family Share",
        "hdoInfo": [
          {
            "hdoServiceCode": "hdo01234567",
            "number": "08011112222",
            "iccid": "DN000011112222",
            "regulation": false,
            "sms": false,
            "voice": false,
            "couponUse": false,
            "coupon": [
              {
                "volume": 42,
                "expire": "201403",
                "type": "bundle"
              }
            ]
          }
        ],
        "coupon": [
          {
            "volume": 42,
            "expire": "201403",
            "type": "bundle"
          }
        ]
      }
    ]
  }
]
```

### Coupon Change
クーポンのON/OFFを切り替えます。クーポン利用状態が変化しないリクエスト(現在 couponUse: true な回線に対する couponUse: true のリクエストなど)の場合もエラーとして判定されず、正常応答(200)が返ります

```
PUT /coupon/
```

#### Required Parameters
<table>
  <tr>
    <th>Name</th>
    <th>Type</th>
    <th>Description</th>
    <th>Example</th>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/couponUse</strong></td>
    <td><em>boolean</em></td>
    <td>クーポン使用中（クーポンON）の場合に true, それ以外は false</td>
    <td><code>false</code></td>
  </tr>
  <tr>
    <td><strong>couponInfo/hdoInfo/hdoServiceCode</strong></td>
    <td><em>string</em></td>
    <td>hdoサービスコード</td>
    <td><code>"hdo01234567"</code></td>
  </tr>
</table>



#### Curl Example
```bash
$ curl -n -X PUT https://api.iijmio.jp/mobile/d/v1/coupon/ \
  -H "Content-Type: application/json" \
  -d '{
  "couponInfo": [
    {
      "hdoInfo": [
        {
          "hdoServiceCode": "hdo01234567",
          "couponUse": false
        }
      ]
    }
  ]
}'

```


#### Response Example
```
HTTP/1.1 200 OK
```
```json
{
  "returnCode": "OK",
  "couponInfo": [
    {
      "hddServiceCode": "hdd01234567",
      "plan": "Family Share",
      "hdoInfo": [
        {
          "hdoServiceCode": "hdo01234567",
          "number": "08011112222",
          "iccid": "DN000011112222",
          "regulation": false,
          "sms": false,
          "voice": false,
          "couponUse": false,
          "coupon": [
            {
              "volume": 42,
              "expire": "201403",
              "type": "bundle"
            }
          ]
        }
      ],
      "coupon": [
        {
          "volume": 42,
          "expire": "201403",
          "type": "bundle"
        }
      ]
    }
  ]
}
```


## Packet
データ利用状況

### Attributes
<table>
  <tr>
    <th>Name</th>
    <th>Type</th>
    <th>Description</th>
    <th>Example</th>
  </tr>
  <tr>
    <td><strong>packetLogInfo/hddServiceCode</strong></td>
    <td><em>string</em></td>
    <td>hddサービスコード</td>
    <td><code>"hdd01234567"</code></td>
  </tr>
  <tr>
    <td><strong>packetLogInfo/hdoInfo/hdoServiceCode</strong></td>
    <td><em>string</em></td>
    <td>hdoサービスコード</td>
    <td><code>"hdo01234567"</code></td>
  </tr>
  <tr>
    <td><strong>packetLogInfo/hdoInfo/packetLog/date</strong></td>
    <td><em>string</em></td>
    <td>通信を行った日付</td>
    <td><code>"201403"</code></td>
  </tr>
  <tr>
    <td><strong>packetLogInfo/hdoInfo/packetLog/withCoupon</strong></td>
    <td><em>integer</em></td>
    <td>クーポン使用状態でのデータ利用量（MB単位）</td>
    <td><code>59</code></td>
  </tr>
  <tr>
    <td><strong>packetLogInfo/hdoInfo/packetLog/withoutCoupon</strong></td>
    <td><em>integer</em></td>
    <td>クーポン不使用状態でのデータ利用量（MB単位）</td>
    <td><code>21</code></td>
  </tr>
  <tr>
    <td><strong>packetLogInfo/plan</strong></td>
    <td><em>string</em></td>
    <td>プラン名<br/><b>one of:</b><code>"Family Share"</code> or <code>"Minimum Start"</code> or <code>"Light Start"</code></td>
    <td><code>"Family Share"</code></td>
  </tr>
  <tr>
    <td><strong>returnCode</strong></td>
    <td><em>string</em></td>
    <td>リターンコード</td>
    <td><code>"OK"</code></td>
  </tr>
</table>

### Packet List
データ利用量を照会します。過去30日分と当日のデータ利用量（MB単位）が日付順に返ります。通信が行われなかった日についても、0 として値が返ります。

```
GET /log/packet/
```


#### Curl Example
```bash
$ curl -n -X GET https://api.iijmio.jp/mobile/d/v1/log/packet/
```


#### Response Example
```
HTTP/1.1 200 OK
```
```json
[
  {
    "returnCode": "OK",
    "packetLogInfo": [
      {
        "hddServiceCode": "hdd01234567",
        "plan": "Family Share",
        "hdoInfo": [
          {
            "hdoServiceCode": "hdo01234567",
            "packetLog": [
              {
                "date": "201403",
                "withCoupon": 59,
                "withoutCoupon": 21
              }
            ]
          }
        ]
      }
    ]
  }
]
```


