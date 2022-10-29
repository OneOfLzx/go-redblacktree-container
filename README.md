GO官方库没有set和自排序的map，平时用起来就差点意思，所以索性想自己实现一个。  
底层基于红黑树实现的容器，包括：set、map。  
底层红黑树是import的我另一个仓库：[go-redblacktree](https://github.com/OneOfLzx/go-redblacktree)  
### 类和接口介绍  
### Set  
`setNodeValueEntry:` 封装底层红黑树节点的数据项，可以存放一个任意类型的数据。实现`RedBlackTreeNodeValEntry`接口  
`setNodeValueEntryComparator:` 用于比较`setNodeValueEntry`之间的大小。实现`NodeValueEntryComparator`接口  
Equal(): 判断两个数据是否相等  
Smaller()：判断数据是否小于  
`SetExtComparatorFunc:` 如果`setNodeValueEntry`中存的不是内置数据类型，则调用者需要实现这个函数来让底层的红黑树知道如何去比较这两个数据项（包括用于相等比较、大小比较）  
  
`SetIterator:` 指向`Set`内部节点的迭代器  
Value(): 返回当前节点的值  
Prev(): 指向前一个节点  
Next(): 指向后一个节点  
  
`Set:` 实际的set数据结构  
NewSet(): 创建一个`Set`。如果准备向`Set`中存的数据不是内置数据类型，则调用者需要传入`SetExtComparatorFunc`函数来让底层的红黑树知道如何去比较这两个数据项（包括用于相等比较、大小比较） ，否则传入nil即可。如果有传入对应函数，且数据是内置数据类型，则优先使用传入的函数来比较  
Find(): 查找一个数据项是否已经在`Set`中了。如果存在返回对应节点的迭代器，否则返回nil  
Insert(): 把一个数据项插入到`Set`中。如果成功返回对应节点的迭代器，否则返回nil  
Remove(): 把一个数据项从`Set`中删除  
RemoveByIter(): 通过迭代器把一个数据项从`Set`中删除  
Size(): 返回`Set`中节点的个数  
### Map  
`mapNodeValueEntry:` 封装底层红黑树节点的数据项，存放一个任意类型的key和存放一个任意类型的value。实现`RedBlackTreeNodeValEntry`接口  
`mapNodeValueEntryComparator:` 用于比较`mapNodeValueEntry`之间的key的大小。实现`NodeValueEntryComparator`接口  
Equal(): 判断两个数据是否相等  
Smaller()：判断数据是否小于  
`MapExtComparatorFunc:` 如果`mapNodeValueEntry`中存的key不是内置数据类型，则调用者需要实现这个函数来让底层的红黑树知道如何去比较这两个数据项（包括用于相等比较、大小比较）  
  
`MapIterator:` 指向`Map`内部节点的迭代器  
Key(): 返回当前节点的key值  
Value(): 返回当前节点的value值  
Prev(): 指向前一个节点  
Next(): 指向后一个节点  
  
`Map:` 实际的map数据结构  
NewMap(): 创建一个`Map`。如果准备向`Map`中存的数据的key值不是内置数据类型，则调用者需要传入`MapExtComparatorFunc`函数来让底层的红黑树知道如何去比较这两个数据项（包括用于相等比较、大小比较） ，否则传入nil即可。如果有传入对应函数，且key值是内置数据类型，则优先使用传入的函数来比较  
Find(): 查找一个key值是否已经在`Map`中了。如果存在返回对应节点的迭代器，否则返回nil  
Insert(): 把一个key-value数据项插入到`Map`中。如果成功返回对应节点的迭代器，否则返回nil。如果key值已经在`Map`中了，则会修改key-value数据项的value值  
Remove(): 把一个key值从`Map`中删除  
RemoveByIter(): 通过迭代器把一个key值从`Map`中删除  
Size(): 返回`Map`中节点的个数