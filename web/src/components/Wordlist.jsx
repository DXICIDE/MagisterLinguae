import { useState, useEffect } from 'react';

function Wordlist({ activeTab }) {
    const [Words, setWords] = useState(null);

    useEffect(() => {
    fetch(`/api/words?language_id=${activeTab.ID}`, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    })
    .then(response => response.json())
    .then(data => setWords(data))
    
    }, [activeTab]);
    
    return (  
        <div>
        <p> Word known frequency lastseen</p>
        {Words?.map((Word, i) => (
            <>
                
                <p key={i}> {Word.word} {Word.known ? "Yes" : "No"}  {Word.frequency}  {Word.last_seen_at}</p>
            </>
            ))}
        </div>
    );
}

export default Wordlist;