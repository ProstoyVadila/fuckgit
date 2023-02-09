import './body.css';
import React from 'react';

let Body = () => {
    const stupidOptions = [
        "как бросить курить",
        "как вернуть бывшую",
        "как вернуть 2007",
        "как попасть в FAANG",
        "как провалить испытательный срок",
        "как не получить оффер",
        "как смотреть в глаза коллегам",
    ];
    const randomOption = Math.floor(Math.random() * stupidOptions.length);

    return (
        <div className="Body">
            <h1>Этот сайт подскажет:</h1>
            <ul class="">
                <li>как исправить содеянное</li>
                <li>как вернуть всё как было</li>
                <li><s>{stupidOptions[randomOption]}</s></li>
            </ul>
        </div>
    );
}

export default Body;
