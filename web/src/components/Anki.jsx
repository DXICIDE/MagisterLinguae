import { useState, useEffect } from 'react';

function Anki({ activeTab }) {
    const [Array, setArray] = useState(null);
    const [index, setIndex] = useState(0);
    const [correct, setCorrect] = useState(0);
    const [wrong, setWrong] = useState(0);
    const [anki, setAnki] = useState(50);
    const [prog, setProg] = useState("start");
    
    function ankistart() {
    fetch(`/api/practice/start`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
    count: anki,
    language_id: activeTab.ID,
    }),
    })
    .then(response => response.json())
    .then(data => setArray(data))
    .then(setIndex(0))
    .then(setCorrect(0))
    .then(setWrong(0))
    .then(setProg("running"))
    }

    useEffect(() => {
    if (Array && index >= Array.length && index > 0) {
        setProg("end");
    }
    }, [index]);

    
    function ankianswer(word, known) {
    fetch(`/api/practice/answer`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
    token_name: word,
    language_id: activeTab.ID,
    known: known
    }),
    })
    .then(response => response.json())
    .then(setIndex(index => index + 1))
    }

    return (  
        <div>
            {prog == "start" && (
            <div>
            <p>Input how many words you would like to practice</p>
            <button onClick={() => ankistart()}>Start</button>
            <input 
            type="number" 
            value={anki} 
            onChange={(e) => setAnki(Number(e.target.value))} 
            min={1}
            max={50}
            />
            </div>
            )}
            {Array && Array.length > 0 && index < anki && prog == "running" && (
            <div>
                <p>{Array[index]?.word}</p>
                <button onClick={() => (ankianswer(Array[index].word, true), setCorrect(correct => correct + 1))}>Know</button>
                <button onClick={() => (ankianswer(Array[index].word, false), setWrong(wrong => wrong + 1))}>Can't remember</button>
            </div>
            )}            
            {Array && Array.length === 0 && (
            <p>All Caught Up!</p>
            )}
            {prog == "end" && (
                <div>
                    <p>Completed!</p>
                    <p>Correct:{correct} Wrong:{wrong}</p>
                    <button onClick={() => setProg("start")}>Practice Again</button>
                </div>
            )}
        </div>
    );
}

export default Anki;