// Date:   Sat Sep 20 03:23:45 PM 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package search_service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/xiaoqixian/v2ex/backend/app/home/util"
	common_util "github.com/xiaoqixian/v2ex/backend/app/common/util"
)

func SearchKeyword(ginCtx *gin.Context) {
	keyword := ginCtx.Query("keyword")
	if keyword == "" {
		ginCtx.JSON(http.StatusBadRequest, gin.H {
			"error": "Keyword is empty",
		})
		return
	}

	es, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{fmt.Sprintf("http://%s:9200", common_util.GetEnv("ESADDR", "localhost"))},
			Username:  "elastic",
			Password:  "changeme",
	})
	if err != nil {
			log.Fatalf("Error creating client: %s", err)
	}

	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"title": "%s"
			}
		}
	}`, keyword)

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("v2ex_posts"),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response: %s", err)
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"err": "Search internal server error",
		})
		return
	}

	defer res.Body.Close()

	var r map[string]any
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	hits := r["hits"].(map[string]any)["hits"].([]any)
	hits_postids := make([]uint64, len(hits))
	for i, hit := range hits {
		log.Printf("  %s\n", hit.(map[string]any)["_source"].(map[string]any)["title"].(string))
		hits_postids[i], err = strconv.ParseUint(hit.(map[string]any)["_id"].(string), 10, 64)
		if err != nil {
			log.Fatalf("Error parsing post id: %s", err)
		}
	}

	posts := util.GetPostsByIds(hits_postids)
	if posts == nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H {
			"error": "Unable to get posts",
		})
		return
	}
	
	ginCtx.JSON(http.StatusOK, gin.H {
		"posts": posts,
	})
}
