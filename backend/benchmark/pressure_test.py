# Date:   Mon Sep 22 04:37:11 PM 2025
# Mail:   lunar_ubuntu@qq.com
# Author: https://github.com/xiaoqixian

import asyncio
import random
import aiohttp

URL_TEMPLATE = "http://192.168.0.10:8000/post/{}"
MAX_POST_ID = 20000   # 根据你的帖子数量设置
NUM_TASKS = 100_000    # 总请求任务数
REQUESTS_PER_TASK = 10
CONCURRENT_LIMIT = 1000  # 同时并发请求数

async def fetch(session, post_id):
    url = URL_TEMPLATE.format(post_id)
    try:
        async with session.get(url, timeout=5) as response:
            print(f"{url} -> {response.status}")
    except Exception as e:
        print(f"{url} -> ERROR: {e}")

async def worker(semaphore):
    async with aiohttp.ClientSession() as session:
        for _ in range(REQUESTS_PER_TASK):
            post_id = random.randint(1, MAX_POST_ID)
            async with semaphore:
                await fetch(session, post_id)

async def main():
    semaphore = asyncio.Semaphore(CONCURRENT_LIMIT)
    tasks = [asyncio.create_task(worker(semaphore)) for _ in range(NUM_TASKS)]
    await asyncio.gather(*tasks)

if __name__ == "__main__":
    asyncio.run(main())
