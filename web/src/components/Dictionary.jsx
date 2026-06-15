import { useState } from 'react';

function Dictionary({ activeTab }) {
    const [inputText, setInputText] = useState("");
    const [Result, setResult] = useState(null);
  
  //api for the text
    function handleDict() {
    if (inputText.trim() === "") {
        setResult(null);
        return;
    }
    fetch(`/api/dictionary/${inputText}?language_id=${activeTab.ID}`, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    })
    .then(response => {
          if (!response.ok) {
            throw new Error("The word was not found")
          }
        return response.json(); 
        }) 
    .then(data => setResult(data))
    .catch(data => setResult({error: "The word was not found"}))
  }
  

    return (
    <div className='dictionary'>
      <div className='dict-container'>
      <h5 className='dict-h5'>Dictionary</h5>
        <textarea className="InputDictionary"
            spellCheck={false}
            value={inputText}
            onChange={(e) => setInputText(e.target.value)} 
            rows={1} 
            cols={5}
            placeholder="Type here..."
        />
        <button className='dict-btn' onClick={handleDict}>Process</button>
        </div>
        <p>{Result?.error}</p>
        {Result && !Result.error && (
        <div className="outputDictionary">
            <h4>{Result?.word}</h4>
            {Result?.entries?.map((entry, i) => (
            <>
              <p key={i}>Part Of Speech: {entry.partOfSpeech}</p>
              {entry.senses?.map((sense, j) => (
              <p key={j}>{j+1}. Definition: {sense.definition} Tag: {sense.tags.join(", ")} {sense.Examples}</p>
              ))}
            </>
            ))}
            <p>Source: <a href={Result?.source?.url}>{Result?.source?.url}</a></p>
            <p>License: {Result?.source?.license?.name} ({Result?.source?.license?.url})</p>
        </div>
        )}
    </div>    
  );
}
  
export default Dictionary;