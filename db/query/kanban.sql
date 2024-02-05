-- name: ListTasks :many
SELECT * FROM kb_task ORDER BY 1;

-- name: GetTasks :many
SELECT * FROM kb_task WHERE task_id = $1 ORDER BY 1;

-- name: CreateTask :one
INSERT INTO kb_task (task_id, title, status, summary, type, priority, tags, estimate, assignee, rankid,
                     color, classname, dependencies, description, due_date, start_date)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16 )
RETURNING *;

-- name: UpdateTask :one
UPDATE kb_task 
SET
    task_id = $1,
    title = $2,  
    status = $3, 
    summary = $4,  
    type = $5,  
    priority = $6,  
    tags = $7,  
    estimate = $8,  
    assignee = $9,  
    rankid = $10, 
    color = $11,  
    classname = $12,  
    dependencies = $13,  
    description = $14,  
    due_date = $15,  
    start_date = $16 
WHERE task_id = $1
RETURNING *;

-- name: ListPriority :many
SELECT * FROM kb_priority ORDER BY 1;

-- name: ListStatus :many
SELECT * FROM kb_status ORDER BY 1;

-- name: ListSubTask :many
SELECT * FROM kb_subtask ORDER BY 1;

-- name: ListTeam :many
SELECT * FROM kb_team ORDER BY 1;

-- name: ListType :many
SELECT * FROM kb_type ORDER BY 1;

