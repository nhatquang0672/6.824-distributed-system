#!/bin/bash
git add .
read time <<< $(echo $(date) | awk '{print $2, $3, $6}')
git commit -m "updated at ""$time"
git push https://'nhatquang0672':'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'@github.com/nhatquang0672/6.824-distributed-system