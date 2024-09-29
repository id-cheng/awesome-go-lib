lux "https://www.youtube.com/watch?v=dQw4w9WgXcQ"  #下载视频
lux -i "https://www.youtube.com/watch?v=dQw4w9WgXcQ" # 该选项显示所有可用的视频质量，而无需下载。-i
lux -f stream "URL"-i # 下载输出中列出的特定流
lux -i -p "https://www.bilibili.com/bangumi/play/ep198061" #下载播放列表
lux -i "https://www.bilibili.com/video/av21877586" "https://www.bilibili.com/video/av21990740" #一次下载多个 URL
lux -i ep198381 av21877586 # 直接用码下载bilibili的视频
lux -i -d "http://www.bilibili.com/video/av20088587" # 调试模式
lux -j "https://www.bilibili.com/video/av20203945" # JSON 格式打印提取的数据
