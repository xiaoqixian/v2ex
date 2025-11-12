# Date:   Mon Sep 22 05:24:46 PM 2025
# Mail:   lunar_ubuntu@qq.com
# Author: https://github.com/xiaoqixian

import asyncio
import aiomysql
import random

DB_HOST = "127.0.0.1"   # MySQL 地址
DB_PORT = 3306
DB_USER = "lunar"
DB_PASSWORD = "lunar"
DB_NAME = "v2ex"

NUM_TASKS = 100       # 并发协程数
QUERIES_PER_TASK = 10  # 每个协程查询次数
POST_ID_MAX = 30000    # 假设 posts 表最大 post_id

async def query_post(pool: aiomysql.Pool, post_id: int):
    async with pool.acquire() as conn:
        async with conn.cursor() as cur:
            await cur.execute("SELECT id, title FROM posts WHERE id=%s", (post_id,))
            result = await cur.fetchone()
            if result:
                print(f"Post {result[0]}: {result[1]}")
            else:
                print(f"Post {post_id} not found")

async def worker(pool: aiomysql.Pool):
    for _ in range(QUERIES_PER_TASK):
        post_id = random.randint(1, POST_ID_MAX)
        try:
            await query_post(pool, post_id)
        except Exception as e:
            print(f"Error querying post {post_id}: {e}")

async def main():
    pool = await aiomysql.create_pool(
        host=DB_HOST,
        port=DB_PORT,
        user=DB_USER,
        password=DB_PASSWORD,
        db=DB_NAME,
        minsize=5,
        maxsize=20,  # 每个 pool 最多连接数
    )

    tasks = [asyncio.create_task(worker(pool)) for _ in range(NUM_TASKS)]
    await asyncio.gather(*tasks)
    pool.close()
    await pool.wait_closed()

if __name__ == "__main__":
    asyncio.run(main())
