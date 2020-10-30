define({ "api": [
  {
    "type": "post",
    "url": "/backend/attr/create",
    "title": "创建属性",
    "name": "CreateAttr",
    "group": "商品属性管理",
    "parameter": {
      "fields": {
        "公共参数": [
          {
            "group": "公共参数",
            "type": "String",
            "optional": true,
            "field": "request_id",
            "description": "<p>请求ID</p>"
          }
        ],
        "请求参数": [
          {
            "group": "请求参数",
            "type": "String",
            "optional": false,
            "field": "Name",
            "description": "<p>属性名称</p>"
          },
          {
            "group": "请求参数",
            "type": "Number",
            "optional": true,
            "field": "FillType",
            "defaultValue": "2",
            "description": "<p>属性值填充类型 [1: 选项框 | 2: 输入框]</p>"
          },
          {
            "group": "请求参数",
            "type": "Number",
            "optional": true,
            "field": "IsRequired",
            "defaultValue": "2",
            "description": "<p>属性是否为必填项 [1:是 | 2: 否]</p>"
          },
          {
            "group": "请求参数",
            "type": "Number",
            "optional": true,
            "field": "IsNumeric",
            "defaultValue": "2",
            "description": "<p>属性值是否为数字类型 [1:是 | 2: 否]</p>"
          },
          {
            "group": "请求参数",
            "type": "String",
            "optional": true,
            "field": "Unit",
            "description": "<p>属性单位</p>"
          },
          {
            "group": "请求参数",
            "type": "Number",
            "optional": true,
            "field": "IsGeneric",
            "defaultValue": "2",
            "description": "<p>是否是SPU通用属性 [1:是 | 2: 否]</p>"
          },
          {
            "group": "请求参数",
            "type": "Number",
            "optional": true,
            "field": "IsSearching",
            "defaultValue": "2",
            "description": "<p>是否用于搜索过滤 [1:是 | 2: 否]</p>"
          },
          {
            "group": "请求参数",
            "type": "String",
            "optional": true,
            "field": "Segments",
            "description": "<p>分割</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "响应参数": [
          {
            "group": "响应参数",
            "type": "String",
            "optional": false,
            "field": "request_id",
            "description": "<p>返回请求的ID.</p>"
          },
          {
            "group": "响应参数",
            "type": "Number",
            "optional": false,
            "field": "code",
            "description": "<p>返回CODE码</p>"
          },
          {
            "group": "响应参数",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>返回信息</p>"
          },
          {
            "group": "响应参数",
            "type": "Object[]",
            "optional": false,
            "field": "data",
            "description": "<p>返回具体内容</p>"
          },
          {
            "group": "响应参数",
            "type": "Number",
            "optional": false,
            "field": "data.attr_id",
            "description": "<p>属性ID</p>"
          },
          {
            "group": "响应参数",
            "type": "String",
            "optional": false,
            "field": "next",
            "description": "<p>&quot;&quot;</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "响应示例:",
          "content": "{\n  \"request_id\": \"c3b51fe3-0dc6-4b69-9a66-508ecc9a8633\",\n  \"code\": 200,\n  \"msg\": \"请求成功\"\n  \"data\": {\n     \"attr_id\": 100031\n   },\n   \"next\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "错误码": [
          {
            "group": "错误码",
            "type": "Number",
            "optional": false,
            "field": "180001",
            "description": "<p>抱歉，创建失败!~</p>"
          },
          {
            "group": "错误码",
            "type": "Number",
            "optional": false,
            "field": "180002",
            "description": "<p>抱歉，删除失败!~</p>"
          },
          {
            "group": "错误码",
            "type": "Number",
            "optional": false,
            "field": "180003",
            "description": "<p>抱歉，删除失败，该属性被其他属性模板引用!~</p>"
          },
          {
            "group": "错误码",
            "type": "Number",
            "optional": false,
            "field": "180004",
            "description": "<p>抱歉，更新失败!~</p>"
          },
          {
            "group": "错误码",
            "type": "Number",
            "optional": false,
            "field": "180005",
            "description": "<p>抱歉，属性名已存在，请换一个名称重试!~</p>"
          },
          {
            "group": "错误码",
            "type": "Number",
            "optional": false,
            "field": "180006",
            "description": "<p>抱歉，抱歉，没有找到相关的属性信息!~</p>"
          },
          {
            "group": "错误码",
            "type": "Number",
            "optional": false,
            "field": "180007",
            "description": "<p>抱歉，由于属性为数值类型，单位为必填项!~</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "异常示例:",
          "content": "{\n  \"request_id\": \"c3b51fe3-0dc6-4b69-9a66-508ecc9a8633\",\n  \"code\": 180001,\n  \"msg\": \"抱歉，创建失败!~\"\n  \"data\": null,\n  \"next\": \"\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "api/backend/attr.go",
    "groupTitle": "商品属性管理"
  }
] });
