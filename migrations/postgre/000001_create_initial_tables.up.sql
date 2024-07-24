
-- Create Users table
CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       registration_date DATE NOT NULL DEFAULT CURRENT_DATE,
                       role VARCHAR(20) NOT NULL CHECK (role IN ('administrator', 'manager', 'developer'))
);

-- Create Projects table
CREATE TABLE projects (
                          id UUID PRIMARY KEY,
                          title VARCHAR(100) NOT NULL,
                          description TEXT,
                          start_date DATE NOT NULL,
                          end_date DATE,
                          manager_id UUID NOT NULL,
                          FOREIGN KEY (manager_id) REFERENCES users(id)
);

-- Create Tasks table
CREATE TABLE tasks (
                       id UUID PRIMARY KEY,
                       title VARCHAR(100) NOT NULL,
                       description TEXT,
                       priority VARCHAR(10) NOT NULL CHECK (priority IN ('low', 'medium', 'high')),
                       status VARCHAR(20) NOT NULL CHECK (status IN ('new', 'in_progress', 'completed')),
                       responsible_id UUID NOT NULL,
                       project_id UUID NOT NULL,
                       creation_date DATE NOT NULL DEFAULT CURRENT_DATE,
                       completion_date DATE,
                       FOREIGN KEY (responsible_id) REFERENCES users(id),
                       FOREIGN KEY (project_id) REFERENCES projects(id)
);



-- Add indexes for foreign keys
CREATE INDEX idx_tasks_responsible_id ON tasks(responsible_id);
CREATE INDEX idx_tasks_project_id ON tasks(project_id);
CREATE INDEX idx_projects_manager_id ON projects(manager_id);