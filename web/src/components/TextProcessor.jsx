import { useState } from 'react';

function TextProcessor({ activeTab }) {
  const [inputText, setInputText] = useState("");
  const [processedResult, setProcessedResult] = useState("");
  
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
    .then(data => setProcessedResult(data.processed_text));
  }

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
        <div className="myDiv">
            {processedResult ? renderProcessedText(processedResult) : "Processed text will appear here..."} 
        </div>
    </div>
    
  );
}


export default TextProcessor;