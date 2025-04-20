import { useState } from 'react'
import { login, setAuthToken } from '../services/auth'
import { useNavigate } from 'react-router-dom'

interface LoginModalProps {
  onClose: () => void
}

const LoginModal = ({ onClose }: LoginModalProps) => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const navigate = useNavigate()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setIsLoading(true)

    console.log('Отправка данных для входа:', { email, password })

    try {
      const response = await login({ email, password })
      console.log('Ответ от сервера:', response)
      setAuthToken(response.token)
      onClose()
      navigate('/dashboard') // Перенаправляем на защищенную страницу
    } catch (err) {
      console.error('Ошибка при входе:', err)
      setError('Неверный email или пароль')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="modal-overlay">
      <div className="modal">
        <button className="modal-close" onClick={onClose}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M18 6L6 18M6 6L18 18" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
          </svg>
        </button>
        <h2 className="modal-title">Вход в систему</h2>
        <p className="modal-subtitle">Введите свои данные для входа</p>
        {error && <p className="error-message">{error}</p>}
        <form onSubmit={handleSubmit} className="form">
          <div className="form-group">
            <input
              type="email"
              placeholder="Почта"
              className="input"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <input
              type="password"
              placeholder="Пароль"
              className="input"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button 
            type="submit" 
            className="button-primary"
            disabled={isLoading}
          >
            {isLoading ? 'Вход...' : 'Войти'}
          </button>
          <button type="button" className="google-button">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M19.8055 10.2303C19.8055 9.55056 19.7491 8.86711 19.6274 8.19836H10.2002V12.0492H15.6015C15.3775 13.2911 14.6572 14.3898 13.6027 15.0879V17.5866H16.8253C18.7176 15.8449 19.8055 13.2728 19.8055 10.2303Z" fill="#4285F4"/>
              <path d="M10.2002 20.0006C12.897 20.0006 15.1714 19.1151 16.8286 17.5865L13.606 15.0879C12.7096 15.6979 11.5521 16.0433 10.2036 16.0433C7.59432 16.0433 5.38268 14.2832 4.58904 11.9169H1.26367V14.4927C2.92252 17.8695 6.36293 20.0006 10.2002 20.0006Z" fill="#34A853"/>
              <path d="M4.58551 11.9169C4.16637 10.6749 4.16637 9.33008 4.58551 8.08811V5.51233H1.26366C-0.154113 8.33798 -0.154113 11.667 1.26366 14.4927L4.58551 11.9169Z" fill="#FBBC04"/>
              <path d="M10.2002 3.95805C11.6256 3.936 13.0035 4.47247 14.036 5.45722L16.8911 2.60218C15.0833 0.904587 12.6838 -0.0287217 10.2002 0.000673888C6.36293 0.000673888 2.92252 2.13185 1.26367 5.51234L4.58551 8.08813C5.37562 5.71811 7.59107 3.95805 10.2002 3.95805Z" fill="#EA4335"/>
            </svg>
            Войти через Google
          </button>
          <a href="#" className="forgot-password">Забыли пароль?</a>
          <p className="terms">
            Нажимая кнопку «Войти», вы соглашаетесь с{' '}
            <a href="#">условиями использования</a>
          </p>
        </form>
      </div>
    </div>
  )
}

export default LoginModal 