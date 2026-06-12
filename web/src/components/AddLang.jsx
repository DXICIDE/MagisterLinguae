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
    <div>
        <textarea className="InputName"
            spellCheck={false}
            value={inputName}
            onChange={(e) => setInputName(e.target.value)} 
            rows={1} 
            cols={5}
            placeholder="Type the name of the tab"
        />
        <textarea className="InputCode"
            spellCheck={false}
            value={inputCode}
            onChange={(e) => setInputCode(e.target.value)} 
            rows={1} 
            cols={2}
            placeholder="Type the international code of the lang"
        />
        <button onClick={() => addLanguage()}>Add</button>
    </div>    
    );
}

export default AddLanguage;