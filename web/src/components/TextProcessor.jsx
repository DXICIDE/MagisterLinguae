import { useState, useEffect } from 'react';
import { FiArrowRight } from 'react-icons/fi';

function TextProcessor({ activeTab }) {
  const [inputText, setInputText] = useState("");
  const [processedResult, setProcessedResult] = useState("");
  const [percentage, setPercentage] = useState(0);
  const [message, setMessage] = useState("");
  const [badgeClass, setbadgeClass] = useState("");

  //api for the text
  function handleProcess(isRefresh = false) {
  fetch('/api/texts/process', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
        text: inputText,
        language_id: activeTab.ID,
        refresh: isRefresh
      })
    })
    .then(response => response.json())
    .then(data => {
      setProcessedResult(data.processed_text);
      calculateDifficulty(data.stats);
    })
  }

  useEffect(() => {
    if (inputText.trim() !== "") {
        handleProcess(false);
    }
  }, [activeTab]);

  function renderProcessedText(text) {
    console.log("Input to render:", text);
    const parts = text.split(/(\[[^\]]+\])/g);
    console.log("Parts:", parts);
    return parts.map((part, index) => {
        if (part.startsWith('[')) {
            const word = part.slice(1, -1);
            return (
                <span key={index} 
                      className="unknown-word" 
                      onClick={() => markWord(word)}>
                    {word}
                </span>
            );
        }
        return <span key={index}>{part}</span>;
    });
  }

  function markWord(word) {
      fetch(`/api/words/${word}/mark?language_id=${activeTab.ID}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    })
    .then(response => response.json())
    .then(() => handleProcess(true))
  }
  
  function calculateDifficulty(stats) {
    const pct = stats.known_words / stats.total_words;
    setPercentage(pct)
    if (pct >= 0.79 && pct <= 0.91) {
      setMessage("Ideal difficulty")
      setbadgeClass("ideal")
    }
    if (0.91 < pct ) {
      setMessage("Too easy!")
      setbadgeClass("too-easy")
    }
    if ( pct < 0.79 ) {
      setMessage("pretty hard")
      setbadgeClass("pretty-hard")
    }

    if ( pct < 0.69 ) {
      setMessage("Too hard!")
      setbadgeClass("too-hard")
    }
    
  }
  
  return (
    <div className="processor-layout">
        <textarea className="input-box"
            spellCheck={false}
            value={inputText} 
            onChange={(e) => setInputText(e.target.value)} 
            rows={15} 
            cols={60}
            placeholder="Paste your text here..."
        />
        <div className="button-container">
          <button className="button-process" onClick={() => handleProcess(false)}>
            <span className="arrow-icon">{'>'}</span>
          </button>
        </div>
        <div className="output-area">
          {processedResult ? renderProcessedText(processedResult) : ""} 
          {processedResult && (
            <span className={`difficulty-badge ${badgeClass}`}>
              {message} {Math.round(percentage * 100)}% known
            </span>
            )}
        </div>
    </div>
  );
}


export default TextProcessor;