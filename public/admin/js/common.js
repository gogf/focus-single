// 全局管理对象
gf = {
    // 刷新验证码
    reloadCaptcha: function() {
        $("img.captcha").attr("src","/captcha?v="+Math.random());
    },
}

// 统一处理Ajax请求成功结果
gf.handleAjaxSuccess = function (r, callback) {
    let options = {
        icon:              "success",
        text:              r.message,
        confirmButtonText: "确定"
    }
    switch (r.code) {
        case  0: options.icon = "success"; break; // 执行成功
        case  1: options.icon = "error";   break; // 执行报错
        case -1: options.icon = "info";    break; // 用户需要登录
    }
    if (typeof options.text == "undefined" || options.text.length == 0 ) {
        switch (options.icon) {
            case "success": options.text = "请求执行成功"; break;
            case "error":   options.text = "请求执行失败"; break;
            case "info":    options.text = "请求执行提示"; break;
        }
    }
    Swal.fire(options).then((value) => {
        if (typeof options.redirect != "undefined" && r.redirect.length > 1) {
            window.location.href = r.redirect
        } else {
            if (typeof callback == "function") {
                callback(r)
            } else {
                // 只有请求执行成功才刷新页面
                switch (options.icon) {
                    case "success": window.location.reload(); break;
                }
            }
        }
    });
}
// 统一处理Ajax请求系统错误
gf.handleAjaxError = function (r, callback) {
    let options = {
        icon:              "error",
        text:              "系统错误",
        confirmButtonText: "确定"
    }
    if (r.responseText != "") {
        options.title = "系统错误"
        options.text  = r.responseText
    }
    Swal.fire(options).then((value) => {
        if (typeof callback == "function") {
            callback(r)
        } else {
            window.location.reload()
        }
    });
}

jQuery(function ($) {
    // 为必填字段添加提示
    $('.required').prepend('<span class="required-mark">*</span>');

    // 初始化select2选择框
    $('.select2').select2();

    $.extend($.validator.messages, {
        required:    "这是必填字段",
        remote:      "请修正此字段",
        email:       "请输入有效的电子邮件地址",
        url:         "请输入有效的网址",
        date:        "请输入有效的日期",
        dateISO:     "请输入有效的日期 (YYYY-MM-DD)",
        number:      "请输入有效的数字",
        digits:      "只能输入数字",
        creditcard:  "请输入有效的信用卡号码",
        equalTo:     "你的输入不相同",
        extension:   "请输入有效的后缀",
        maxlength:   $.validator.format("最多可以输入 {0} 个字符"),
        minlength:   $.validator.format("最少要输入 {0} 个字符"),
        rangelength: $.validator.format("请输入长度在 {0} 到 {1} 之间的字符串"),
        range:       $.validator.format("请输入范围在 {0} 到 {1} 之间的数值"),
        max:         $.validator.format("请输入不大于 {0} 的数值"),
        min:         $.validator.format("请输入不小于 {0} 的数值")
    });
})