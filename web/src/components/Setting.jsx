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
    <div>
        <button key={"delete"} onClick={() => (deleteTab())}> <p>Delete current language</p> </button>
        <button key={"reset"} onClick={() => (resetLang(), setIsOpen(true))}> <p>Reset current language</p> </button>
        {isOpen && (
        <div className="modal-overlay">
            <div className="modal-content">
                <h3>Language Reset</h3>
                <p>{deleteResponse.message}, words deleted: {deleteResponse.words_deleted}</p>
                <button onClick={() => setIsOpen(false)}>OK</button>
            </div>
        </div>
        )}
    </div>
    );
}

export default Settings;