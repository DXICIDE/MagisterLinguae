import { useState, useEffect } from 'react';

function TextProcessor({ activeTab }) {
  const [inputText, setInputText] = useState("");
  const [processedResult, setProcessedResult] = useState("");
  const [percentage, setPercentage] = useState(0);
  const [message, setMessage] = useState("");

  //api for the text
  function handleProcess() {
  fetch('/api/texts/process', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
        text: inputText,
        language_id: activeTab.ID
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
        handleProcess();
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
    .then(() => handleProcess())
  }
  
  function calculateDifficulty(stats) {
    const pct = stats.known_words / stats.total_words;
    setPercentage(pct)
    if (pct >= 0.79 && pct <= 0.91) {
      setMessage("Ideal difficulty")
    }
    if (0.91 < pct ) {
      setMessage("Too easy!")
    }
    if ( pct < 0.79 ) {
      setMessage("Pretty hard")
    }

    if ( pct < 0.69 ) {
      setMessage("Too hard!")
    }
    
  }
  
  return (
    <div>
        <textarea className="InputBox"
            spellCheck={false}
            value={inputText} 
            onChange={(e) => setInputText(e.target.value)} 
            rows={15} 
            cols={60}
            placeholder="Paste your text here..."
        />
        <button onClick={handleProcess}>Process</button>
        <p>{processedResult ? `${message}: ${percentage*100}% known` : ""}</p>
        <div className="myDiv">
            {processedResult ? renderProcessedText(processedResult) : "Processed text will appear here..."} 
        </div>
    </div>
    
  );
}


export default TextProcessor;