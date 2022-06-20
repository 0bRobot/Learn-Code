#define _CRT_SECURE_NO_WARNINGS
#include <stdio.h>
#include <stdlib.h>
#include <graphics.h>

int main()
{
	initgraph(600 ,880);
	IMAGE  imgs;  //声明存储图片的变量
	//loadimage (&imgs,"123.png");  //改为宽字节     图片加载到内存
	loadimage(&imgs, "123.png",560,600);//输出是压缩图片  让图片适用窗口
	// 560,300, 显示图片的大小
	// ,0,300   图片坐标从哪个位置开始  原图显示下半部分
	//putimage(0, 0, 560, 300, &imgs, 0, 300);   // 图片的输出  及位置   原图显示下半部分
	putimage(0,0, 560,300, &imgs,0,0);   //显示上半部分
	//思考 让图片 慢慢显示完成    for 循环

	//鼠标操作
	MOUSEMSG m;  // 声明一个存储鼠标消息的变量
	while (1)  //一直获取鼠标消息
	{
		m = GetMouseMsg(); //获取鼠标消息
		switch (m.uMsg)
		{
		case WM_MOUSEMOVE: //鼠标移动

			break;
		case WM_LBUTTONDOWN: //左键鼠标按下
			circle(m.x,m.y,50);  // 画圆 以鼠标坐标为基准
			break;
		case WM_RBUTTONDOWN: //鼠标邮件按下
			rectangle(m.x-5,m.y-5,m.x+10,m.y+10);
			break;
		}
	}



	getchar();
	closegraph();
	return 0;
}