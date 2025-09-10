import React, { useState } from "react";

const TaskForm = ({ addTask }) => {
  const [title, setTitle] = useState("");
  const [dueDate, setDueDate] = useState("");
  const [priority, setPriority] = useState("medium");

  const handleSubmit = (e) => {
    e.preventDefault();
    addTask(title, dueDate, priority);
    setTitle("");
    setDueDate("");
    setPriority("medium");
  };

  return (
    <form onSubmit={handleSubmit} className="mb-4 flex flex-col gap-2">
      <input
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Enter task title"
        className="p-2 border rounded"
      />
      <input
        type="datetime-local"
        value={dueDate}
        onChange={(e) => setDueDate(e.target.value)}
        className="p-2 border rounded"
      />
      <select
        value={priority}
        onChange={(e) => setPriority(e.target.value)}
        className="p-2 border rounded"
      >
        <option value="low">Low</option>
        <option value="medium">Medium</option>
        <option value="high">High</option>
      </select>
      <button type="submit" className="p-2 bg-green-500 text-white rounded">
        Add Task
      </button>
    </form>
  );
};

export default TaskForm;
