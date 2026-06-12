import { useState, useEffect } from 'react';
import TextProcessor from './components/TextProcessor';
import Dictionary from './components/Dictionary';
import Wordlist from './components/Wordlist';
import Anki from './components/Anki';
import { Description, Dialog, DialogPanel, DialogTitle } from '@headlessui/react';
import AddLanguage from './components/AddLang';
import Settings from './components/Setting';

function App() {
  const [activeTab, switchTab] = useState(null);
  const [languagesApi, setLang] = useState([]);
  const [activeSection, setActiveSection] = useState("text");
  const [deleteResponse, setDeleteResponse] = useState("");
  
  useEffect(() => {
      refreshLanguages()
      
      fetch('/api/languages/current')
        .then(response => response.json())  
        .then(data => switchTab(data));
      }, []);

function refreshLanguages() {
    fetch('/api/languages')
        .then(response => response.json())
        .then(data => setLang(data));
}


  function handleTabChange(lang) {
    fetch('/api/languages/current', {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ last_language: lang.ID })
    })
    .then(switchTab(lang));
  }
  
  return (
    <div>
      <h1>MagisterLinguae</h1>
      <p>Current language: {activeTab?.Name}</p>
      <p>ID: {activeTab?.ID}</p>
      <p>code: {activeTab?.Code}</p>

      <button key="settings" onClick={() => setActiveSection("settings")}>
          Settings
        </button>
      {activeSection === "settings" && <Settings activeTab={activeTab} onSuccess={() => {
            setActiveSection("text");
            refreshLanguages();
        }} />}
      {languagesApi.map(lang => (
        <button key={lang.ID} onClick={() => handleTabChange(lang)}>
          {lang.Name}
        </button>
      ))}
      <button key={"+"} onClick={() => setActiveSection("+")}>
          {"+"}
      </button>
      <button key={"text"} onClick={() => setActiveSection("text")}>
          {"text"}
      </button>
      <button key={"Wordlist"} onClick={() => setActiveSection("Wordlist")}>
          {"Wordlist"}
      </button>
      <button key={"anki"} onClick={() => setActiveSection("anki")}>
          {"anki"}
      </button>
      {activeSection === "text" && <TextProcessor activeTab={activeTab} />}
      {activeSection === "text" && <Dictionary activeTab={activeTab} />}
      {activeSection === "Wordlist" && <Wordlist activeTab={activeTab} />}
      {activeSection === "anki" && <Anki activeTab={activeTab} />}
      {activeSection === "+" && <AddLanguage onSuccess={() => {
            setActiveSection("text");
            refreshLanguages();
        }} />}
    </div>
  );
}

export default App;