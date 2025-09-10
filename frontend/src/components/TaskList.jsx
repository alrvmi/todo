import React from "react";

const TaskList = ({ tasks, deleteTask, toggleCompletion }) => {
  const activeTasks = tasks.filter((task) => !task.Completed);
  const completedTasks = tasks.filter((task) => task.Completed);

  return (
    <div>
      <h2 className="text-xl font-semibold mb-2">Active Tasks</h2>
      {activeTasks.map((task) => (
        <div
          key={task.ID}
          className="flex items-center justify-between p-2 border-b"
        >
          <div>
            <input
              type="checkbox"
              checked={task.Completed}
              onChange={() => toggleCompletion(task.ID)}
            />
            <span className="ml-2">{task.Title}</span>
            {task.DueDate && (
              <p className="text-sm text-gray-500">{task.DueDate}</p>
            )}
            <p className="text-sm text-gray-500">Priority: {task.Priority}</p>
          </div>
          <button
            onClick={() => deleteTask(task.ID)}
            className="p-1 bg-red-500 text-white rounded"
          >
            Delete
          </button>
        </div>
      ))}
      <h2 className="text-xl font-semibold mt-4 mb-2">Completed Tasks</h2>
      {completedTasks.map((task) => (
        <div
          key={task.ID}
          className="flex items-center justify-between p-2 border-b"
        >
          <div>
            <input
              type="checkbox"
              checked={task.Completed}
              onChange={() => toggleCompletion(task.ID)}
            />
            <span className="ml-2 line-through">{task.Title}</span>
            {task.DueDate && (
              <p className="text-sm text-gray-500">{task.DueDate}</p>
            )}
            <p className="text-sm text-gray-500">Priority: {task.Priority}</p>
          </div>
          <button
            onClick={() => deleteTask(task.ID)}
            className="p-1 bg-red-500 text-white rounded"
          >
            Delete
          </button>
        </div>
      ))}
    </div>
  );
};

export default TaskList;
