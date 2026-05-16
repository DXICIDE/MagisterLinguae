import { useState } from 'react';

function TextProcessor() {
  const [inputText, setInputText] = useState("");
  const [processedResult, setProcessedResult] = useState("");
  
  
  return (
    <div>
        <textarea className="InputBox"
            value={inputText} 
            onChange={(e) => setInputText(e.target.value)} 
            rows={15} 
            cols={60}
            placeholder="Paste your text here..."
        />
        <button>Process</button>
        <div className="myDiv">
             <p>{processedResult || "Processed text will appear here..."}</p>
        </div>
    </div>
    
  );
}

export default TextProcessor;