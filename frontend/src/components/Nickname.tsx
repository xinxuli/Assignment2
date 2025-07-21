import { useState } from 'react';

type NicknameProps = { nickname: string };

function Nickname({ nickname }: NicknameProps) {
  return (
    <div>
      <h2>Hi, {nickname}!</h2>
    </div>
  );
}

export default Nickname;