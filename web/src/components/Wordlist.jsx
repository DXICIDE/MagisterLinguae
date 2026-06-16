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
        <div className='list-container'>
        <div className='list-header'>
            <p className='list-word'>Word</p>
            <p className='list-known'>Known</p>
            <p className='list-frequency'>Freq</p>
            <p className='list-seen'>Last Seen</p>
        </div>
        <div className='list'>
        {Words?.map((Word, i) => (
            <div className='list-items'>
                <p className='list-word'> {Word.word} </p>
                <p className='list-known'> {Word.known ? "Yes" : "No"} </p>
                <p className='list-frequency'> {Word.frequency}  </p>
                <p className='list-seen'> {new Date(Word.last_seen_at).toLocaleDateString('en-GB', { day: 'numeric', month: 'short', year: 'numeric' })}</p>
            </div>
            ))}
        </div>
        </div>
    );
}

export default Wordlist;