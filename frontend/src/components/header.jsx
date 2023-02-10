import React from "react";
import { Link } from "react-router-dom";

const Header = () => {
    return (
        <div className="Header">
            <h1 className="titile">Fuck Git</h1>
            <p className="title title-description">Мы оба знаем, зачем ты тут. Так чего тянуть?</p>
        <Link to="/questions">
            <button className="btn help-me">
                Помоги мне!
            </button>
        </Link>

        </div>
    );
}

export default Header;