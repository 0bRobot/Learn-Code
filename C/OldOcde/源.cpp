#define _CRT_SECURE_NO_WARNINGS
#include<stdio.h>
#include <stdlib.h>
#include <graphics.h>  // 引用三方图形库头文件   

void outtextxy_int(int x, int y, char *format, int iNumber)
{
	char str[20] = "";
	sprintf(str,format,iNumber);
	outtextxy(x,y,str);
}

int main()
{
	initgraph(400, 600);  // 初始化窗口 (int  width,int height )
	//左上角为 原点  0   ，0
	// 横向 为  x 轴
	//纵向 为 y 轴
	setbkcolor(LIGHTRED);//设置背景颜色
	cleardevice();  // 设置背景颜色需要刷新
	

	putpixel(340,268,RED);  //画点    颜色宏 表示颜色  BULE  BLACK 。。。
	for (int x = 0; x < 400; x += 2)  //for 循环画多点 
	{
		 putpixel(x, 268, RED);
	}
	setlinecolor(LIGHTBLUE);  //设置线的颜色
	line(10,10,10,100); //划线
	//画矩形
	setfillcolor(LIGHTCYAN); //设置填充颜色
	rectangle(0, 0, 398, 598);
	fillrectangle(50, 50, 100, 130);
	
	setlinecolor(RGB(54,154,75));  //设置圆 线条颜色
	circle(70,40,40);  //画圆

	setfillcolor(RED);
	solidcircle(100, 210, 34); //无边界线的圆
	
	settextcolor(LIGHTGREEN);
	char outst[] = "NI Hao C!!!!";
	setbkmode(TRANSPARENT); //去掉文字背景
	//for (int i = 0; i <= 200; i += 20)
	//{
	//	cleardevice(); //清除显示设备
	//	settextstyle(i, 0, "Consolas");  //字体逐渐 变大
	//	outtextxy(10, 60, outst);  //文字输出  使用宽字节
	//	Sleep(270);
	//}

	settextstyle(50, 0, "Consolas");  //字体逐渐 变大
	outtextxy(80, 350, outst);  //文字输出  使用宽字节

	outtextxy_int(200, 200, "你好!% d", 1314); //数字输出

	getchar(); //从窗口获取按键  防止窗口关闭

	closegraph(); //关闭窗口
	return 0;
}