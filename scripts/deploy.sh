#!/bin/bash

## 以系统服务后台运行方式管理Golang程序
## 停止服务
## sudo systemctl stop [serviceName]
## 查看服务状态
## sudo systemctl status [serviceName]
## 启动服务
## sudo systemctl start [serviceName]
## 重启服务
## sudo systemctl restart [serviceName]

# 要后台运行的程序
PROGRAM_PATH='/webapp/app/web-app'
# 程序工作目录
PROGRAM_WORK_DIR='/webapp/app'
# 服务名
SERVICE_NAME='Go-Web.service'

# 部署，以系统后台服务方式运行程序.
function deploy_service() {
    # 检查系统中是否已经存在这个服务
    if systemctl list-units --type=service | grep -q "$SERVICE_NAME"; then
        echo "服务 [$SERVICE_NAME] 该服务已经存在"
        exit 1
    else
        echo "开始创建 [$SERVICE_NAME] 服务..."
        # 创建 systemd 服务文件
        cat <<EOF | sudo tee /etc/systemd/system/$SERVICE_NAME > /dev/null
[Unit]
Description=$SERVICE_NAME service file
# 确保程序在网络服务之后启动
After=network.target

[Service]
# Go 程序的可执行文件路径
ExecStart=$PROGRAM_PATH
# 程序的工作目录
WorkingDirectory=$PROGRAM_WORK_DIR
# 程序崩溃后自动重启
Restart=always
# 使用普通用户权限
User=root
# 使用普通用户组
Group=root
# Environment=PATH=/usr/bin:/usr/local/bin:/root/dev/go/bin  # -- 设置环境变量
[Install]
WantedBy=multi-user.target  # -- 在多用户模式下启动服务
EOF
        echo "[$SERVICE_NAME] 服务文件创建成功."
        echo "重新加载 systemd 配置..."
        sudo systemctl daemon-reload

        sudo systemctl start $SERVICE_NAME # 启动系统服务
        echo "启动 [@SERVICE_NAME] 服务成功."
        sudo systemctl enable $SERVICE_NAME
        echo "启动 [$SERVICE_NAME] 服务开机自启动成功."
    fi
}

case "$1" in
    start)
      echo "start [$SERVICE_NAME] Service"
      sudo systemctl start $SERVICE_NAME
      ;;
    stop)
      echo "stop [$SERVICE_NAME] Service"
      sudo systemctl stop $SERVICE_NAME
      ;;
    restart)
      echo "restart [$SERVICE_NAME] Service"
      sudo systemctl restart $SERVICE_NAME
      ;;
    status)
      echo "lookOver [$SERVICE_NAME] Service Status"
      sudo systemctl status $SERVICE_NAME
      ;;
    deploy)
      echo "deploy [$SERVICE_NAME] Service"
      deploy_service
      ;;
    clear)
      echo "clear [$SERVICE_NAME] Service"
      sudo systemctl stop $SERVICE_NAME # 停止服务
      sudo systemctl disable $SERVICE_NAME # 关闭自启服务
      sudo rm -f /etc/systemd/system/$SERVICE_NAME # 删除服务文件
      sudo systemctl daemon-reload # 重新加载缓存
      sudo systemctl reset-failed # 清除服务启动失败的记录
      ;;
esac


