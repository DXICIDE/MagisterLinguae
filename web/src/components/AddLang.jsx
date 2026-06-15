import { useState } from 'react';

function AddLanguage({ onSuccess }) {
    const [inputName, setInputName] = useState("");
    const [inputCode, setInputCode] = useState("");
    const [Result, setResponse] = useState(null);

  //api for the text
    function addLanguage() {
    fetch(`/api/languages`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
    code: inputCode,
    name: inputName,
    }),
    })
    .then(response => response.json())
    .then(data => {
        setResponse(data);
        onSuccess();
    })
    }

    return (
    <div className="add-lang-form">
        <input 
            className="add-lang-input"
            value={inputName}
            onChange={(e) => setInputName(e.target.value)}
            placeholder="Language name"
        />
        <input 
            className="add-lang-input"
            value={inputCode}
            onChange={(e) => setInputCode(e.target.value)}
            placeholder="Language code (e.g., fr)"
        />
        <button className="add-lang-btn" onClick={() => addLanguage()}>Add</button>
    </div>    
    );
}

export default AddLanguage;