define({ "api": [
  {
    "type": "post",
    "url": "/sms_spike",
    "title": "添加秒杀活动",
    "version": "1.0.0",
    "name": "CreateSmsGoods",
    "group": "SmsGoods",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "goods_id",
            "description": "<p>商品id</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "price",
            "description": "<p>商品价格</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "stock",
            "description": "<p>库存</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "start_time",
            "description": "<p>活动开始时间戳</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "end_time",
            "description": "<p>活动结束时间戳</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"goods_id\": 1,\n    \"price\": 500,\n    \"stock\": 100,\n    \"start_time\": 1615784400230000000,\n    \"end_time\": 1615791600230000000\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n \"code\": 1,\n\t\"data\": null,\n\t\"msg\": \"成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "app/controller/sms_spike.go",
    "groupTitle": "SmsGoods"
  }
] });
