import { useState } from 'react';

function Settings({ activeTab, onSuccess }) {
    const [isOpen, setIsOpen] = useState(false);
    const [id, setID] = useState(0)
    const [deleteResponse, setDeleteResponse] = useState("");

    function resetLang() {
    if (!window.confirm("Are you sure?")) {
        return;
    }
    fetch(`/api/db/reset?language_id=${activeTab.ID}`, {
        method: 'POST',
    })
    .then(response => response.json())
    .then(data => {
    setDeleteResponse(data);
    setIsOpen(true)
    });
    }

    function deleteTab() {
    resetLang()
    fetch(`/api/languages/${activeTab.ID}`, {
        method: 'DELETE',
    })
    .then(response => response.json())
    .then(data => {
    setDeleteResponse(data);
    onSuccess();
    });
    }
    
    return (
    <div className='settings'>
        <button className='setting-lang-btn' onClick={() => (deleteTab())}> <p>Delete current language</p> </button>
        <button className='setting-lang-btn' onClick={() => (resetLang())}> <p>Reset current language</p> </button>
        <button className='close-btn' onClick={() => (onSuccess())}> <p>Close</p> </button>
        {isOpen && (
            <div className="Reset-content">
                <p>{deleteResponse.message}, words deleted: {deleteResponse.words_deleted}</p>
                <button className="ok-btn" onClick={() => (onSuccess(), setIsOpen(false))}>OK</button>
            </div>
        )}
    </div>
    );
}

export default Settings;