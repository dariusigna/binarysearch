import React, { useState, useEffect } from 'react';
import './App.css';
import { findIndex } from './apiService';
import { AxiosError } from 'axios';

const App: React.FC = () => {
  const [valueToFind, setValueToFind] = useState<string>("");
  const [index, setIndex] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleFindIndex = async () => {
    try {
      const index = await findIndex(Number(valueToFind));
      setIndex(index);
      setError(null); // Clear any previous error
    } catch (err) {
      if (err instanceof AxiosError) {
        setError(err.response?.data.error || 'An error occurred');
      } else {
        setError('Error finding index. Please try again.');
      }
    }
  };

  return (
      <div className="container">
        <h1 className="title">Index API</h1>
        <div>
          <input
              type="number"
              className="input"
              value={valueToFind}
              onChange={(e) => setValueToFind((e.target.value))}
          />
          <button className="button" onClick={handleFindIndex}>Find Index</button>
          {index !== null && error === null && <p className="index-display">{index}</p>}
          {error && <p className="error-message">{error}</p>}
        </div>
      </div>
  );
};

export default App;