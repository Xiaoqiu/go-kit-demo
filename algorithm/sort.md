#### 快速排序
> 算法描述：是对插入算法的一种优化，利用对问题的二分化，实现递归完成快速排序 ，
在所有算法中二分化是最常用的方式，将问题尽量的分成两种情况加以分析， 最终以形成类似树的方式加以利用，
因为在比较模型中的算法中，最快的排序时间 负载度为 O(nlgn).

##### 算法步骤
- 将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于
- 将两边数据进行递归调用步骤1
- 将所有数据合并

#### 堆排序算法
> 算法描述：首先建一个堆，然后调整堆，调整过程是将节点和子节点进行比较，
> 将 其中最大的值变为父节点，递归调整调整次数lgn,最后将根节点和尾节点交换再n次 调整O(nlgn).

##### 算法步骤
- 创建最大堆或者最小堆（我是最小堆）
- 调整堆
- 交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)
> 堆树的定义：[原文链接](https://blog.csdn.net/guoweimelon/article/details/50904346)
> 堆树的定义如下：

（1）堆树是一颗完全二叉树；

（2）堆树中某个节点的值总是不大于或不小于其孩子节点的值；

（3）堆树中每个节点的子树都是堆树。

当父节点的键值总是大于或等于任何一个子节点的键值时为最大堆。 
当父节点的键值总是小于或等于任何一个子节点的键值时为最小堆。


