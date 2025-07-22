import { userInfoStore } from "../store";

type NicknameProps = { nickname: string };

export const Nickname = () => {
  return (
    <div>
      <h2>Hi, {userInfoStore.nickname}!</h2>
    </div>
  );
}