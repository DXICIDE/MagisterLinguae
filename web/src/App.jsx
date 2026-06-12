import { useState, useEffect } from 'react';
import TextProcessor from './components/TextProcessor';
import Dictionary from './components/Dictionary';
import Wordlist from './components/Wordlist';
import Anki from './components/Anki';
import { Description, Dialog, DialogPanel, DialogTitle } from '@headlessui/react';
import AddLanguage from './components/AddLang';
import Settings from './components/Setting';
import { FiType, FiBook, FiList, FiZap, FiSettings, FiEdit } from 'react-icons/fi';

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
    <div className="app-layout">
      <div className="sidebar">
        <button 
        className={`sidebar-btn ${activeSection === "text" ? "active" : ""}`}
        onClick={() => setActiveSection("text")}>
        <span className="sidebar-btn-content">
          <FiEdit size={20} />
          Learn
        </span>
        </button>
        <button
        className={`sidebar-btn ${activeSection === "Wordlist" ? "active" : ""}`}
        onClick={() => setActiveSection("Wordlist")}>
        <span className="sidebar-btn-content">
          <FiList size={20} />
        Word list
        </span>
        </button>
        <button
        className={`sidebar-btn ${activeSection === "anki" ? "active" : ""}`}
        onClick={() => setActiveSection("anki")}>
        <span className="sidebar-btn-content">
          <FiZap size={20} />
          Anki
        </span>
        </button>
      </div>
      <div className="main-content">

      <h1>MagisterLinguae</h1>
      <p>Current language: {activeTab?.Name}</p>

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
      
      {activeSection === "text" && <TextProcessor activeTab={activeTab} />}
      {activeSection === "text" && <Dictionary activeTab={activeTab} />}
      {activeSection === "Wordlist" && <Wordlist activeTab={activeTab} />}
      {activeSection === "anki" && <Anki activeTab={activeTab} />}
      {activeSection === "+" && <AddLanguage onSuccess={() => {
            setActiveSection("text");
            refreshLanguages();
        }} />}
      </div>
    </div>
  );
}

export default App;