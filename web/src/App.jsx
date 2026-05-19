import { useState, useEffect } from 'react';
import TextProcessor from './components/TextProcessor';

function App() {
  const [activeTab, switchTab] = useState(null);
  const [languagesApi, setLang] = useState([]);

  
  useEffect(() => {
      fetch('/api/languages')
        .then(response => response.json())  
        .then(data => setLang(data));
      console.log("Component loaded!");
      
      fetch('/api/languages/current')
        .then(response => response.json())  
        .then(data => switchTab(data));
      }, []);

  function handleTabChange(lang) {
    fetch('/api/languages/current', {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ last_language: lang.ID })
    })
    .then(() => switchTab(lang));
  }
  
  return (
    <div>
      <h1>MagisterLinguae</h1>
      <p>Current language: {activeTab?.Name}</p>
      
      {languagesApi.map(lang => (
        <button key={lang.ID} onClick={() => handleTabChange(lang)}>
          {lang.Name}
        </button>
      ))}
      <TextProcessor activeTab={activeTab} />
    </div>
  );
}

export default App;