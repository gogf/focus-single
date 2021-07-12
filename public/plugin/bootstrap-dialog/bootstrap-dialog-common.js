// 提示框
dialog = {
    Alert: function (message, handler) {
        handler = handler || null;
        BootstrapDialog.alert({
            title: '提示',
            message: message,
            callback: function () {
                if (handler != null) dialog.f(handler);
            }
        });
    },
    //成功信息SucceedInfo
    Succeed: function (message, handler) {
        handler = handler || null;
        BootstrapDialog.show({
            type: BootstrapDialog.TYPE_SUCCESS,
            title: '信息',
            message: message,
            buttons: [{
                label: '确认',
                action: function (dialog) {
                    if (handler != null) dialog.f(handler);
                    dialog.close();
                }
            }]
        });
    },
    //失败信息ErrorInfo
    Error: function (message, handler) {
        handler = handler || null;
        BootstrapDialog.show({
            type: BootstrapDialog.TYPE_DANGER,
            title: '信息',
            message: message,
            buttons: [{
                label: '确认',
                action: function (dialog) {
                    if (handler != null) dialog.f(handler);
                    dialog.close();
                }
            }]
        });
    },
    //询问信息
    Confirm: function (message, ok_fun, cancel_fun) {
        BootstrapDialog.confirm(message, function (result) {
            if (result) {
                dialog.f(ok_fun);
            } else {
                if (cancel_fun != null) {
                    dialog.f(cancel_fun);
                }
            }
        });
    },
    EmptyFunc: function () {
        return;
    },
    f: function (fn) {
        fn();
    }
}


