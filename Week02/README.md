## Homeworl problems
### 1. dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么?
不应该warp，可以直接处理完ErrNoRows返回nil 或者在service里增加Is判定是否为nil






