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
    <div>
        <textarea className="InputDictionary"Ze
            spellCheck={false}
            value={inputText}
            onChange={(e) => setInputText(e.target.value)} 
            rows={5} 
            cols={5}
            placeholder="Paste your word here..."
        />
        <button onClick={handleDict}>Process</button>
        <p>{Result?.error}</p>
        {Result && !Result.error && (
        <div className="outputDictionary">
            <h4>{Result?.word}</h4>
            {Result?.entries?.map((entry, i) => (
            <>
              <p key={i}>Part Of Speech: {entry.partOfSpeech}</p>
              {entry.senses?.map((sense, j) => (
              <p key={j}>{j+1}. Definition: {sense.definition} Tag: {sense.tags} {sense.Examples}</p>
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

// fmt.Printf("Word: %s\n", responseObject.Word)
// 	for _, entry := range responseObject.Entries {
// 		fmt.Printf("PartOfSpeech: %s\n", entry.PartOfSpeech)
// 		for _, sense := range entry.Senses {
// 			fmt.Printf(" Definition: %s\n", sense.Definition)
// 			for _, tag := range sense.Tags {
// 				fmt.Printf("  Tags: %s\n", tag)-*
// 			}
// 			for _, example := range sense.Examples {
// 				fmt.Printf("  Example: %s\n", example)
// 			}
// 		}
// 	}

  //fmt.Printf("\nSource: %s\n", responseObject.Source.URL)
	// fmt.Println("License by:")
	// fmt.Printf("name: %s\n", responseObject.Source.License.Name)
	// fmt.Printf("url: %s\n\n", responseObject.Source.License.URL)
  
export default Dictionary;