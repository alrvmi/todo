import React from "react";

const FilterSort = ({ setFilter, setSort }) => {
  return (
    <div className="mb-4 flex gap-4">
      <select
        onChange={(e) => setFilter(e.target.value)}
        className="p-2 border rounded"
      >
        <option value="all">All</option>
        <option value="active">Active</option>
        <option value="completed">Completed</option>
      </select>
      <select
        onChange={(e) => setSort(e.target.value)}
        className="p-2 border rounded"
      >
        <option value="date">Date</option>
        <option value="priority">Priority</option>
      </select>
      <select
        onChange={(e) => setFilter(e.target.value)}
        className="p-2 border rounded"
      >
        <option value="">No Date Filter</option>
        <option value="today">Today</option>
        <option value="week">This Week</option>
        <option value="overdue">Overdue</option>
      </select>
    </div>
  );
};

export default FilterSort;
