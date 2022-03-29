package main

/**
冒泡排序
比较相邻的两个元素,不满足大小要求时,交换位置
当某次比较时,没有需要交换的元素,说明数据已经有序,结束
*/
func bubbleSort(nums []int) {

}

/**
插入排序
分为已排序区间和未排序区间,开始时,已排序区间只有一个元素.
遍历未排序区间的元素,在已排序区间找到合适的位置插入
*/
func insertionSort(nums []int) {

}

/**
选择排序
分为已排序区间和未排序区间.
在未排序区间的元素中,找到一个最小的,放入已排序区间的末尾
*/
func selectionSort(nums []int) {

}

/**
归并排序
把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起
*/
func mergeSort(nums []int) {

}

/**
快速排序
选择一个元素作为pivot（分区点）,一般选择数组的最后一个元素.优化可以用随机数选择pivot,将其交换到数组的最后.
遍历数据,将小于 pivot 的放到左边，将大于 pivot 的放到右边.
递归排序左右两部分,知道区间缩小为1
*/
func quicksort(nums []int) {

}

/**
桶排序
将要排序的数据分到几个有序的桶里，每个桶里的数据再单独进行排序。
桶内排完序之后，再把每个桶里的数据按照顺序依次取出，组成的序列就是有序的了。
限制: 1.需要很容易就能划分成 m 个桶，并且，桶与桶之间有着天然的大小顺序。
	2.数据在各个桶之间的分布是比较均匀的

比较适合用在外部排序,如数据量太大,不能全部加载到内存.先遍历数据划分到多个文件(桶)中,每个桶分别加载到内存中排序
*/
func bucketSort(nums []int) {

}

/**
计数排序
桶排序的一种特殊情况
数据范围不大时,记录每个值出现的次数
优化:使用前缀和可以做到稳定排序
*/
func countingSort(nums []int) {

}

/**
基数排序
分为最高位优先和最低位优先
先按照最后一位来排序手机号码，然后，再按照倒数第二位重新排序
按照每位来排序的排序算法要是稳定的，否则这个实现思路就是不正确的。
可以把所有的单词补齐到相同长度，位数不够的可以在后面补“0”
*/
func radixSort(nums []int) {

}

/**
堆排序
*/
func heapSort(nums []int) {

}
