#define _CRT_SECURE_NO_WARNINGS
#include <stdio.h>
#include <conio.h>
#include <stdlib.h>
#include <graphics.h>
#include <mmsystem.h>  //多媒体接口的头文件
#pragma comment(lib,"winmm.lib")//多媒体接口的头文件

#define SNAKE_MUN 1000  //蛇的最大长度
#define W_HIGHT 480  //窗口高
#define W_WIDE  880   //窗口宽
/*
项目 贪吃蛇 

知识点:结构体，循环，函数 ，EasyX [图形界面库]  、 数组
图形化  创建一个图形窗口
*/

enum DIR  //枚举
{
	UP,
	DOWN,
	LEFT,
	RIGHT,
};

struct Snakes  //蛇的结构    
{
	int Size; //蛇的长度
	int Dir;  // 蛇的方向
	int speed; //蛇的速度
	POINT Coor[SNAKE_MUN];// 坐标  x   y  POINT 2 个long整形的 结构体  [] 里面需要数组


}Snakes;

struct Food
{
	int x;   //食物坐标
	int y;
	int r;   //食物大小
	bool flag;//是否被吃
	DWORD color; //食物颜色
}Food;

void GameInit()
{
	//播放背景音乐
	//mciSendString("open ./res/snake_bgm.mp3 alias BGM", 0, 0, 0);
	//mciSendString("play BGM repeat", 0, 0, 0);
	// init 初始化  graph 图形窗口
	initgraph(W_WIDE, W_HIGHT);  //初始化图形窗口并设置大小。(640, 480,SHOWCONSOLE) 显示控制台
	//设置随机数种子  srand()
	srand(GetTickCount());//一般用时间戳  GetTickCount()获取系统开机到现在的毫秒数
								 
	//初始化蛇
	Snakes.Size = 3; //初始大小
	Snakes.speed = 10; //初始速度
	Snakes.Dir = RIGHT;    //方向
	for (int i = 0; i < Snakes.Size; i++)
	{
		Snakes.Coor[i].x = 40-10*i;  //初始y坐标  y坐标向下
		Snakes.Coor[i].y = 10;  //初始x坐标  x坐标向右  左上角 坐标为 0 0 
	}
	//初始化食物
	Food.x = rand()% W_WIDE; //食物的生成 在窗口里面
	//rand() 随机产生一个整数，如果没有设置随机种子，每次都是固定的整数。
	//设置随着种子需要头文件 stdlib.h
	Food.y = rand() % W_HIGHT;
	Food.color = RGB(rand() % 256, rand() % 256, rand() % 256);  //随机颜色
	Food.r = rand() % 10 + 5;
	Food.flag = true;

}

void BKColor()
{
	//双缓冲绘图
	BeginBatchDraw();
	//设置窗口背景
	setbkcolor(RGB(95, 211, 234)); //设置窗口背景颜色
	cleardevice();//设置颜色2部    清空绘图设备
	
	//绘制蛇
	setfillcolor(GREEN);
	for (int i = 0; i < Snakes.Size; i++)
	{
		solidcircle(Snakes.Coor[i].x, Snakes.Coor[i].y, 6);   //绘制一个圆  半径为5
	}

	//绘制食物
	if (Food.flag)
	{
		solidcircle(Food.x, Food.y, Food.r);
	}
	//双缓冲结束
	EndBatchDraw();
}
	
//移动蛇
void MoveSnake()
{
	//移动是坐标发生改变
	//让身体跟着头移动
	for (int i = Snakes.Size-1; i > 0; i--)
	{
		Snakes.Coor[i] = Snakes.Coor[i - 1];
	}	
	switch (Snakes.Dir)
	{
	case UP:
		Snakes.Coor[0].y-=Snakes.speed;
		if (Snakes.Coor[0].y <= 0)  //y 坐标小于0  超出了上面边界
		{
			Snakes.Coor[0].y = W_HIGHT-10;  // 穿墙
		}
		break;
	case DOWN:
		Snakes.Coor[0].y += Snakes.speed;
		if (Snakes.Coor[0].y >= W_HIGHT) //y 坐标大于等于窗口高度  超出了下面面边界
		{
			Snakes.Coor[0].y = 10;
		}
		
		break;
	case LEFT:
		Snakes.Coor[0].x -= Snakes.speed;
		if (Snakes.Coor[0].x <= 0)
		{
			Snakes.Coor[0].x = W_WIDE-10;
		}
		break;
	case RIGHT:
		Snakes.Coor[0].x+=Snakes.speed;
		if (Snakes.Coor[0].x >= W_WIDE)
		{
			Snakes.Coor[0].x = 10;
		}
		break;

	}
}

//通过按键改变蛇的方向
void InputKey()
{
	//判断是否有按键
	if (_kbhit())
	{
		//有按键返回真
		switch (_getch())  //_getch()  获取按键  需要头文件  #include <conio.h>
		//_getch()  阻塞函数  不按键盘会停止程序
		{
			//72  80 75 77  上下左右 键值
		case'w':
		case'W':
		case 72:
			if (Snakes.Dir != DOWN) //控制蛇掉头
			{
				Snakes.Dir = UP;
			}
			break;
		case's':
		case'S':
		case 80:
			if (Snakes.Dir != UP)
			{
				Snakes.Dir = DOWN;
			}
			break;
		case'a':
		case'A':
		case 75:
			if (Snakes.Dir != RIGHT)
			{
				Snakes.Dir = LEFT;
			}
			break;
		case'd':
		case'D':
		case 77:
			if(Snakes.Dir != LEFT)
			{
				Snakes.Dir = RIGHT;
			}
			break;
		case ' ':  //空格  游戏暂停
			while (1)
			{
				if (_getch() == ' ')
					return;
			}
			break;
		}

	}
	
}
void EatFood()
{
	if ( Food.flag && Snakes.Coor[0].x >= Food.x-Food.r && Snakes.Coor[0].x <= Food.x+Food.r &&
		Snakes.Coor[0].y >= Food.y - Food.r && Snakes.Coor[0].y <= Food.y + Food.r)
		//判断蛇是否碰到了 食物坐标。 且 Food flag 为真
	{
		Food.flag = false;  //吃掉以后  flag为 false
		Snakes.Size++;
	}

	if (!Food.flag) //食物消失重新初始化食物
	{
		//初始化食物
		Food.x = rand() % W_WIDE-10; //食物的生成 在窗口里面
		//rand() 随机产生一个整数，如果没有设置随机种子，每次都是固定的整数。
		//设置随着种子需要头文件 stdlib.h
		Food.y = rand() % W_HIGHT-10;
		Food.color = RGB(rand() % 256, rand() % 256, rand() % 256);  //随机颜色
		Food.r = rand() % 10 + 10;
		Food.flag = true;

	}

}

int main()
{
	GameInit();
	
	while (1)//防止图形化退出
	{
		BKColor();
		MoveSnake();
		InputKey();
		EatFood();
		Sleep(100);
	}
	return 0;
}
