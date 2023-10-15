'use client'

import React, { useState, useEffect, useRef } from 'react';
import styles from "./page.module.css";
import { FaLongArrowAltRight } from 'react-icons/fa';

const words = ["Unique", "Developer", "Innovative", "Creative"];
const wordChangeInterval = 3000;
const animationDuration = 1000;

export default function Home() {
    const [wordIndex, setWordIndex] = useState(0);
    const wordRef = useRef<HTMLSpanElement | null>(null);

    useEffect(() => {
        const intervalId = setInterval(() => {
            setWordIndex((prevIndex) => (prevIndex + 1) % words.length);
        }, wordChangeInterval);

        return () => clearInterval(intervalId);
    }, []);

    useEffect(() => {
        if (wordRef.current) {
            wordRef.current.classList.add(styles.wordAnimation);

            setTimeout(() => {
                if (wordRef.current) {
                    wordRef.current.classList.remove(styles.wordAnimation);
                }
            }, animationDuration);
        }
    }, [wordIndex]);

    return (
        <main>
            <div className={styles.top}>
                <div className={styles.centered}>
                    <h1 className={styles.top_text}>
                        Create <span className={styles.accentText}>
                            <span ref={wordRef}>{words[wordIndex]}</span>
                        </span> Profiles
                    </h1>
                    <p className={styles.description_text}>
                        Simplify Your Developer Journey with easy to make and Tailored Profiles on devbio.me
                        <br />
                        Elevating the Developer Experience has never been easier.
                    </p>

                    <button className={styles.get_started_button}>
                        Get Started <FaLongArrowAltRight className={styles.arrow_icon} />
                    </button>

                    <div className={styles.image_div}>
                        <a href="/dashboard">
                            <img src="https://cdn.discordapp.com/attachments/1124442686192095345/1163103137297617017/Profile_Showcase3.png" alt="Example image of a profile made with devbio.me" />
                        </a>
                    </div>
                </div>
            </div>
        </main>
    );
}