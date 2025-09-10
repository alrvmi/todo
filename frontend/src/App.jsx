import React, { useState, useEffect } from "react";
import TaskForm from "./components/TaskForm";
import TaskList from "./components/TaskList";
import FilterSort from "./components/FilterSort";

const App = () => {
  const [tasks, setTasks] = useState([]);
  const [theme, setTheme] = useState("light");
  const [filter, setFilter] = useState("all");
  const [sort, setSort] = useState("date");
  const [error, setError] = useState("");

  useEffect(() => {
    fetchTasks();
  }, [filter, sort]);

  const fetchTasks = async () => {
    try {
      const filteredTasks = await window.go.main.App.FilterTasks(filter, "");
      const sortedTasks = await window.go.main.App.SortTasks(sort);
      setTasks(sortedTasks);
    } catch (err) {
      setError("Failed to load tasks");
    }
  };

  const addTask = async (title, dueDate, priority) => {
    if (!title) {
      setError("Task title cannot be empty");
      return;
    }
    try {
      await window.go.main.App.AddTask(title, dueDate, priority);
      fetchTasks();
      setError("");
    } catch (err) {
      setError("Failed to add task");
    }
  };

  const deleteTask = async (id) => {
    if (window.confirm("Are you sure you want to delete this task?")) {
      try {
        await window.go.main.App.DeleteTask(id);
        fetchTasks();
      } catch (err) {
        setError("Failed to delete task");
      }
    }
  };

  const toggleCompletion = async (id) => {
    try {
      await window.go.main.App.ToggleTaskCompletion(id);
      fetchTasks();
    } catch (err) {
      setError("Failed to update task");
    }
  };

  const toggleTheme = () => {
    setTheme(theme === "light" ? "dark" : "light");
  };

  return (
    <div
      className={`min-h-screen ${
        theme === "light" ? "bg-gray-100" : "bg-gray-900 text-white"
      } p-4`}
    >
      <div className="max-w-2xl mx-auto">
        <h1 className="text-2xl font-bold mb-4">To-Do List</h1>
        <button
          onClick={toggleTheme}
          className="mb-4 p-2 bg-blue-500 text-white rounded"
        >
          Toggle {theme === "light" ? "Dark" : "Light"} Theme
        </button>
        {error && <p className="text-red-500">{error}</p>}
        <TaskForm addTask={addTask} />
        <FilterSort setFilter={setFilter} setSort={setSort} />
        <TaskList
          tasks={tasks}
          deleteTask={deleteTask}
          toggleCompletion={toggleCompletion}
        />
      </div>
    </div>
  );
};

export default App;
