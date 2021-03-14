define({ "api": [
  {
    "type": "post",
    "url": "/spike/buy",
    "title": "秒杀购买商品",
    "version": "1.0.0",
    "name": "Buy",
    "group": "Spike",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "user_id",
            "description": "<p>用户id</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "spike_id",
            "description": "<p>秒杀活动id</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n \"user_id\": 432,\n \"spike_id\": 1\n}",
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
    "filename": "app/controller/buy.go",
    "groupTitle": "Spike"
  },
  {
    "type": "post",
    "url": "/spike/cache",
    "title": "缓存秒杀商品库存",
    "version": "1.0.0",
    "name": "CacheSpike",
    "group": "Spike",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>秒杀活动id</p>"
          },
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
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "redis_key",
            "description": "<p>Redis队列键名，代表库存</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"id\": 12,\n    \"goods_id\": 2,\n    \"price\": 500,\n    \"stock\": 100,\n    \"start_time\": 1615784400230000000,\n    \"end_time\": 1615791600230000000,\n    \"redis_key\": \"spike-stock-12\"\n}",
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
    "filename": "app/controller/buy.go",
    "groupTitle": "Spike"
  }
] });