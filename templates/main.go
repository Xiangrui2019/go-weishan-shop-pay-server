package templates

var WeixinNotifySenderTemplate = `商品名称: %s
商品ID: %s
收货人姓名: %s
收货地址: %s
收货人电话号码: %s
订单备注: %s
购买数量: %v
付款价格:￥%v
<br />
<a href='%s/%s'>1. 检查是否已经发货</a>
<br />
<a href='%s/map.html?addr=%s'>2. 查看地图</a>
<br />
<a href='%s/%s'>3. 确认发货</a>`
