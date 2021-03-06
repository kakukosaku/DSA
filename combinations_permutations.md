# 排列组合实现

概述，排列组合大家高中都有学习过A_m^n/C_m^n分别表示排列组合。这里要说的是如何用python实现排列组合

**组合**
这里以Python官方给出的生成器实现为例说明。[传送门](https://docs.python.org/3.6/library/itertools.html#itertools.combinations)

```python3
def combinations(iterable, r):
    # combinations('ABCD', 2) --> AB AC AD BC BD CD
    # combinations(range(4), 3) --> 012 013 023 123
    pool = tuple(iterable)
    n = len(pool)
    if r > n:
        return
    indices = list(range(r))
    yield tuple(pool[i] for i in indices)
    while True:
        for i in reversed(range(r)):
            if indices[i] != i + n - r:
                break
        else:
            return
        indices[i] += 1
        for j in range(i+1, r):
            indices[j] = indices[j-1] + 1
        yield tuple(pool[i] for i in indices)
```

这段代码，我初看很好奇。生成器，我比较熟悉了；组合的数学定义我也很清楚。但究竟要如何才能实现代码示例中给出的效果？之前写one-line形式的[逆波兰表达式](https://zh.wikipedia.org/wiki/%E9%80%86%E6%B3%A2%E5%85%B0%E8%A1%A8%E7%A4%BA%E6%B3%95)时收获的经验告诉我，写程序针对一类情形时，应先**“发现”规律**！

**这段代码核心问题在于**
1. 由可迭代对对象'ABCD'给出AB AC AD BC BD CD的返回值，无非生成其索引indices的无序集合，即01 02 03 12 13 23(与给定值r相关)

2. 使实现的更优雅

先说第一个问题：找规律，考虑更大范围上的集合，更容易总结出规律(盲目穷举，在小集合上可以正确穷举所有组合，但依照某种方法，才能更适用，不遗漏！)

**算法部分**
考虑：

['A', 'B', 'C', 'D', 'E', 'F']  欲取其3个元素的组合等同于  
[ 0,   1,   2,   3,   4,   5]

如前所述，以下面这样的”规律“找出其组合

[0, 1, 2]

[0, 1, 3]

[0, 1, 4]

[0, 1, 5] **注意1**

[0, 2, 3] **注意2**

[0, 2, 4]

[0, 2, 5] **注意1**

[0, 3, 4] **注意2**

[0, 3, 5]

[0, 4, 5] **注意3**

[1, 2, 3]**注意2**

[1, 2, 4]
.
.
.
[3, 4, 5] **终点！**

上面一串序列是以一样的原则，列举了组合的全部情形(无遗漏，重复较盲目穷举)，在我标记出注意1,2,3的地方，有如下解释：


注意1的目的在于：对初始序列[0, 1, 2](与r相关)，对于正在被递增的元素(最后一个为例)，何时为递增终点？

注意2的目的在于：对初始序列[0 ,1, 2]达到递增终点后的处理？

注意3的目的在于：对其它被递增的元素，何时为递增终点？

假设:
- 总元素数为n(例中为6)
- 被求组合范围为r(例中为3)

在每一位(i)的终点判定中

如果i没有达到终点，对它做递增处理

到达终点后，(下一位未到终点情况下)由下一位的值+1得到i的重新循环值

**这个过程中，主要的循环在于最后一个元素的递增，以及最后一个元素到终点后，继续前一个循环的初值！**

总感觉我说的不够明白，但你看想清楚了吗？

**实现部分**

官方给出的代码，实现已经很简洁了。如果生成器部分我不再赘述，需要说明的for-else结构。

代码中使用for-else结构，进行如果i没有达到终点时的处理。即for 中所有循环执行完，没有提前break，执行return。意味着所有元素都已到达终点。