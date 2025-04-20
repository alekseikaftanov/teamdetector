interface LoginModalProps {
  onClose: () => void
}

export const LoginModal = ({ onClose }: LoginModalProps) => {
  return (
    <div className="modal-overlay">
      <div className="modal">
        <button className="modal-close" onClick={onClose}>
          ✕
        </button>
        <h2 className="modal-title">Вход</h2>
        <p className="modal-subtitle">
          Войдите, чтобы заказать бесплатное демо —<br />
          так мы сможем с вами связаться :)
        </p>
        <form className="form">
          <div className="form-group">
            <input
              type="email"
              className="input"
              placeholder="Почта"
            />
          </div>
          <div className="form-group">
            <input
              type="password"
              className="input"
              placeholder="Пароль"
            />
            <a href="/forgot-password" className="forgot-password">
              Забыли пароль?
            </a>
          </div>
          <button type="submit" className="button-secondary">
            Войти
          </button>
          <button type="button" className="google-button">
            <img src="/google.svg" alt="" />
            Войти с помощью Google
          </button>
        </form>
        <p className="terms">
          При входе вы принимаете <a href="/terms">какой-нибудь документ</a><br />
          и <a href="/privacy">ещё какой-нибудь документ</a>
        </p>
      </div>
    </div>
  )
} 