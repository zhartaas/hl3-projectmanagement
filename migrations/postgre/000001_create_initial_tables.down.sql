-- down.sql

-- Drop indexes
DROP INDEX IF EXISTS idx_tasks_responsible_id;
DROP INDEX IF EXISTS idx_tasks_project_id;
DROP INDEX IF EXISTS idx_projects_manager_id;

-- Drop Tables
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS users;