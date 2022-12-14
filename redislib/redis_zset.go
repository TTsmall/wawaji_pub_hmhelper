package redislib

import "github.com/garyburd/redigo/redis"

//ZUNIONSTORE 	不实现
//ZINTERSTORE

/*
ZADD key score member [[score member] [score member] ...]
将一个或多个 member 元素及其 score 值加入到有序集 key 当中。
如果某个 member 已经是有序集的成员，那么更新这个 member 的 score 值，并通过重新插入这个 member 元素，来保证该 member 在正确的位置上。
score 值可以是整数值或双精度浮点数。
如果 key 不存在，则创建一个空的有序集并执行 ZADD 操作。
当 key 存在但不是有序集类型时，返回一个错误。
对有序集的更多介绍请参见 sorted set 。
时间复杂度:
O(M*log(N))， N 是有序集的基数， M 为成功添加的新成员的数量。
返回值:
被成功添加的新成员的数量，不包括那些被更新的、已经存在的成员。
*/
func (rd *RedisHandleModel) Zadd(key string, score interface{}, member string) (int, error) {
	return redis.Int(rd.Do("ZADD", key, score, member))
}

//看Zadd
func (rd *RedisHandleModel) Zadds(key string, li map[string]interface{}) (int, error) {
	p := []interface{}{
		key,
	}
	for k, v := range li {
		p = append(p, v, k)
	}
	return redis.Int(rd.Do("ZADD", p...))
}

/*
ZCARD key
返回有序集 key 的基数。
返回值:
当 key 存在且是有序集类型时，返回有序集的基数。
当 key 不存在时，返回 0 。
*/
func (rd *RedisHandleModel) Zcard(key string) (int, error) {
	return redis.Int(rd.Do("ZCARD", key))
}

/*
ZCOUNT key min max
返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量。
区间及无限
min 和 max 可以是 -inf 和 +inf ，这样一来，
你就可以在不知道有序集的最低和最高 score 值的情况下，使用 ZRANGEBYSCORE 这类命令。
默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于)。
时间复杂度:
O(log(N)+M)， N 为有序集的基数， M 为值在 min 和 max 之间的元素的数量。
返回值:
score 值在 min 和 max 之间的成员的数量。
*/
func (rd *RedisHandleModel) Zcount(key string, min, max interface{}) (int, error) {
	return redis.Int(rd.Do("ZCOUNT", key, min, max))
}

/*
ZINCRBY key increment member
为有序集 key 的成员 member 的 score 值加上增量 increment 。
可以通过传递一个负数值 increment ，让 score 减去相应的值，比如 ZINCRBY key -5 member ，就是让 member 的 score 值减去 5 。
当 key 不存在，或 member 不是 key 的成员时， ZINCRBY key increment member 等同于 ZADD key increment member 。
当 key 不是有序集类型时，返回一个错误。
score 值可以是整数值或双精度浮点数。
时间复杂度:
O(log(N))
返回值:
member 成员的新 score 值，以字符串形式表示。
*/
func (rd *RedisHandleModel) Zincrby(key string, inc float64, member string) (string, error) {
	return redis.String(rd.Do("ZINCRBY", key, inc, member))
}

/*
ZRANGE key start stop [WITHSCORES]
返回有序集 key 中，指定区间内的成员。
其中成员的位置按 score 值递增(从小到大)来排序。
具有相同 score 值的成员按字典序(lexicographical order )来排列。
如果你需要成员按 score 值递减(从大到小)来排列，请使用 ZREVRANGE 命令。
下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。
你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。
超出范围的下标并不会引起错误。
比如说，当 start 的值比有序集的最大下标还要大，或是 start > stop 时， ZRANGE 命令只是简单地返回一个空列表。
另一方面，假如 stop 参数的值比有序集的最大下标还要大，那么 Redis 将 stop 当作最大下标来处理。
可以通过使用 WITHSCORES 选项，来让成员和它的 score 值一并返回，返回列表以 value1,score1, ..., valueN,scoreN 的格式表示。
客户端库可能会返回一些更复杂的数据类型，比如数组、元组等。
返回值:
指定区间内，成员列表。
*/
func (rd *RedisHandleModel) Zrange(key string, start, stop int) ([]string, error) {
	return redis.Strings(rd.Do("ZRANGE", key, start, stop))
}

/*
看Zrange
返回值:
指定区间内，带有 score 值(可选)的有序集成员的列表。
*/
func (rd *RedisHandleModel) ZrangeWithScores(key string, start, stop int) (map[string]string, error) {
	return redis.StringMap(rd.Do("ZRANGE", key, start, stop, "WITHSCORES"))
}

/*
看Zrange
返回值:
指定区间内，带有 score 值(可选)的有序集成员的列表。
*/
func (rd *RedisHandleModel) ZrangeWithScoresByInt(key string, start, stop int) (map[string]int, error) {
	return redis.IntMap(rd.Do("ZRANGE", key, start, stop, "WITHSCORES"))
}

/*
ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。
有序集成员按 score 值递增(从小到大)次序排列。
时间复杂度:
O(log(N)+M)， N 为有序集的基数， M 为被结果集的基数。
返回值:
指定区间内，带有 score 值(可选)的有序集成员的列表。
区间及无限
min 和 max 可以是 -inf 和 +inf ，这样一来，你就可以在不知道有序集的最低和最高 score 值的情况下，使用 ZRANGEBYSCORE 这类命令。
默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于)。
举个例子：
ZRANGEBYSCORE zset (1 5
返回所有符合条件 1 < score <= 5 的成员，而
ZRANGEBYSCORE zset (5 (10
则返回所有符合条件 5 < score < 10 的成员。
*/
func (rd *RedisHandleModel) ZrangeByScore(key string, min, max interface{}) ([]string, error) {
	return redis.Strings(rd.Do("ZRANGEBYSCORE", key, min, max))
}

//看ZrangeByScore
func (rd *RedisHandleModel) ZrangeByScoreLimit(key string, min, max interface{}, offset, count int) ([]string, error) {
	return redis.Strings(rd.Do("ZRANGEBYSCORE", key, min, max, offset, count))
}

//看ZrangeByScore
func (rd *RedisHandleModel) ZrangeByScoreWithScore(key string, min, max interface{}) (map[string]string, error) {
	return redis.StringMap(rd.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES"))
}

//看ZrangeByScore
func (rd *RedisHandleModel) ZrangeByScoreWithScoreLimit(key string, min, max interface{}, offset, count int) (map[string]string, error) {
	return redis.StringMap(rd.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES", offset, count))
}

/*
ZRANK key member
返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递增(从小到大)顺序排列。
排名以 0 为底，也就是说， score 值最小的成员排名为 0 。
使用 ZREVRANK 命令可以获得成员按 score 值递减(从大到小)排列的排名。
时间复杂度:
O(log(N))
返回值:
如果 member 是有序集 key 的成员，返回 member 的排名。
如果 member 不是有序集 key 的成员，返回 nil 。
*/
func (rd *RedisHandleModel) Zrank(key, member string) (int, error) {
	return redis.Int(rd.Do("ZRANK", key, member))
}

/*
ZREM key member [member ...]
移除有序集 key 中的一个或多个成员，不存在的成员将被忽略。
当 key 存在但不是有序集类型时，返回一个错误。
时间复杂度:
O(M*log(N))， N 为有序集的基数， M 为被成功移除的成员的数量。
返回值:
被成功移除的成员的数量，不包括被忽略的成员。
*/
func (rd *RedisHandleModel) Zrem(key string, members ...string) (int, error) {
	p := []interface{}{
		key,
	}
	for _, v := range members {
		p = append(p, v)
	}
	return redis.Int(rd.Do("ZREM", p...))
}

/*
ZREMRANGEBYRANK key start stop
移除有序集 key 中，指定排名(rank)区间内的所有成员。
区间分别以下标参数 start 和 stop 指出，包含 start 和 stop 在内。
下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。
你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。
时间复杂度:
O(log(N)+M)， N 为有序集的基数，而 M 为被移除成员的数量。
返回值:
被移除成员的数量。
*/
func (rd *RedisHandleModel) ZremRangeByRank(key string, start, stop int) (int, error) {
	return redis.Int(rd.Do("ZREMRANGEBYRANK", key, start, stop))
}

/*
ZREMRANGEBYSCORE key min max
移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。
自版本2.1.6开始， score 值等于 min 或 max 的成员也可以不包括在内，详情请参见 ZRANGEBYSCORE 命令。
时间复杂度:
O(log(N)+M)， N 为有序集的基数，而 M 为被移除成员的数量。
返回值:
被移除成员的数量。
*/
func (rd *RedisHandleModel) ZremRangeByScore(key string, min, max interface{}) (int, error) {
	return redis.Int(rd.Do("ZREMRANGEBYSCORE", key, min, max))
}

/*
ZREVRANGE key start stop [WITHSCORES]
返回有序集 key 中，指定区间内的成员。
其中成员的位置按 score 值递减(从大到小)来排列。
具有相同 score 值的成员按字典序的逆序(reverse lexicographical order)排列。
除了成员按 score 值递减的次序排列这一点外， ZREVRANGE 命令的其他方面和 ZRANGE 命令一样。
时间复杂度:
O(log(N)+M)， N 为有序集的基数，而 M 为结果集的基数。
返回值:
指定区间内，带有 score 值(可选)的有序集成员的列表。
*/
func (rd *RedisHandleModel) ZrevRange(key string, start, stop int) ([]string, error) {
	return redis.Strings(rd.Do("ZREVRANGE", key, start, stop))
}

/*
看ZrevRange
返回值:
指定区间内，带有 score 值(可选)的有序集成员的列表。
*/
func (rd *RedisHandleModel) ZrevRangeWithScores(key string, start, stop int) (map[string]string, error) {
	return redis.StringMap(rd.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
}

/*
看ZrevRange
返回值:
指定区间内，带有 score 值(可选)的有序集成员的列表。
*/
func (rd *RedisHandleModel) ZrevRangeWithScoresByInt(key string, start, stop int) (map[string]int, error) {
	return redis.IntMap(rd.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
}

/*
ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
返回有序集 key 中， score 值介于 max 和 min 之间(默认包括等于 max 或 min )的所有的成员。有序集成员按 score 值递减(从大到小)的次序排列。
具有相同 score 值的成员按字典序的逆序(reverse lexicographical order )排列。
除了成员按 score 值递减的次序排列这一点外， ZREVRANGEBYSCORE 命令的其他方面和 ZRANGEBYSCORE 命令一样。
时间复杂度:
O(log(N)+M)， N 为有序集的基数， M 为结果集的基数。
返回值:
指定区间内，带有 score 值(可选)的有序集成员的列表。
*/
func (rd *RedisHandleModel) ZrevRangeByScore(key string, min, max interface{}) ([]string, error) {
	return redis.Strings(rd.Do("ZREVRANGEBYSCORE", key, min, max))
}

//看ZrevRangeByScore
func (rd *RedisHandleModel) ZrevRangeByScoreLimit(key string, min, max interface{}, offset, count int) ([]string, error) {
	return redis.Strings(rd.Do("ZREVRANGEBYSCORE", key, min, max, offset, count))
}

//看ZrevRangeByScore
func (rd *RedisHandleModel) ZrevRangeByScoreWithScore(key string, min, max interface{}) (map[string]string, error) {
	return redis.StringMap(rd.Do("ZREVRANGEBYSCORE", key, min, max, "WITHSCORES"))
}

//看ZrevRangeByScore
func (rd *RedisHandleModel) ZrevRangeByScoreWithScoreLimit(key string, min, max interface{}, offset, count int) (map[string]string, error) {
	return redis.StringMap(rd.Do("ZREVRANGEBYSCORE", key, min, max, "WITHSCORES", offset, count))
}

/*
ZREVRANK key member
返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递减(从大到小)排序。
排名以 0 为底，也就是说， score 值最大的成员排名为 0 。
使用 ZRANK 命令可以获得成员按 score 值递增(从小到大)排列的排名。
时间复杂度:
O(log(N))
返回值:
如果 member 是有序集 key 的成员，返回 member 的排名。
如果 member 不是有序集 key 的成员，返回 nil 。
*/
func (rd *RedisHandleModel) ZrevRank(key, member string) (int, error) {
	return redis.Int(rd.Do("ZREVRANK", key, member))
}

/*
ZSCORE key member
返回有序集 key 中，成员 member 的 score 值。
如果 member 元素不是有序集 key 的成员，或 key 不存在，返回 nil 。
返回值:
member 成员的 score 值，以字符串形式表示。
*/
func (rd *RedisHandleModel) Zscore(key, member string) (string, error) {
	return redis.String(rd.Do("ZSCORE", key, member))
}
