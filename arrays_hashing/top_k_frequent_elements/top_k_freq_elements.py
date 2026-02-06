def top_k_freq_elements(nums, k):
    count = {}
    for n in nums:
        # set the count to 1 but also add the existing count for n.
        count[n] = 1 + count.get(n, 0)

    buckets = [[] for _ in range(len(nums) + 1)]
    for num, cnt in count.items():
        buckets[cnt].append(num)

    result = []
    for i in range(len(buckets) - 1, 0, -1):
        for n in buckets[i]:
            result.append(n)
            if len(result) == k:
                return result

    return result
